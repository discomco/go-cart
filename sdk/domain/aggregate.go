package domain

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/model"
)

const (
	aggregateStartVersion                = -1 // used for EventStoreDB
	aggregateAppliedEventsInitialCap     = 10
	aggregateUncommittedEventsInitialCap = 10
)

type AggregateSetter interface {
	SetAggregate(a IAggregate)
}

type EvtTypeGetter interface {
	GetEventType() EventType
}

type CmdTypeGetter interface {
	GetCommandType() CommandType
}

//IAggPlugin (Aggregate Plugin) is an injector that allows us to inject ITryCmd and IApplyEvt injectors into the Aggregate in an elegant way.
type IAggPlugin interface {
	AggregateSetter
}

// IApplyEvt is an IAggPlugin injector that allows us to inject Event Appliers into the Aggregate
type IApplyEvt interface {
	IAggPlugin
	EvtTypeGetter
	ApplyEvent(evt IEvt, state model.IWriteModel) error
}

// ITryCmd is an IAggPlugin injector that allows us to inject Command Executors into the Aggregate
type ITryCmd interface {
	IAggPlugin
	CmdTypeGetter
	TryCommand(ctx context.Context, command ICmd) (IEvt, dtos.IFbk)
}

// IAggregate is the injector for Aggregates
// In an Event Sourced application, the Aggregate can be considered as the equivalent of the ActiveRecord in classic CrUD (Create, Upsert, Delete) applications.
// The Aggregate can be considered as the heart of an ES system, that unites State with Behavior.
// Its main responsibilities are:
//  1. to build a (volatile) State from an ordered list of previously committed Events, that are sourced from the Event Stream that is identified by the Id.
//  2. CanAcceptName Command requests and applying business logic that checks whether the Command (ICmd) is allowed to be executed or not,
//     according to a number of Specifications, the Current State (see 1.) and the Command's Payload
//  3. If Command execution is allowed, Raise a new Event and ApplyEvent it to itself, as to update the Current State to the New State.
// Specifically for GO-SCREAM CMD Applications, given their modular nature, we rely on Aggregate Composition.
// Aggregate Composition is a technique that allows us to inject
// a series of AggPluginFtor functors that create a feature's IApplyEvt and ITryCmd injectors,
// in order to compose an Aggregate that has all the capabilities required to process the Event Stream.
type IAggregate interface {
	ISimpleAggregate
	Inject(agg IAggregate, actors ...AggPluginFtor) IAggregate
	KnowsCmd(topic CommandType) bool
	KnowsEvt(topic EventType) bool
}

// ISimpleAggregate is an injector that abstracts the implementation of a traditional monolithic Aggregate.
type ISimpleAggregate interface {
	AggregateTypeGetter
	String() string
	TryCommand(ctx context.Context, command ICmd) (IEvt, dtos.IFbk)
	ApplyEvent(event IEvt, isCommitted bool) error
	GetState() model.IWriteModel
	GetUncommittedEvents() []IEvt
	ClearUncommittedEvents()
	SetAppliedEvents(events []IEvt)
	GetAppliedEvents() []IEvt
	ToSnapshot()
	RaiseEvent(event IEvt) error
	GetVersion() int64
	GetID() core.IIdentity
	SetID(identity core.IIdentity) IAggregate
}

type AggFtor func() IAggregate
type AggBuilder func() IAggregate
type AggPluginFtor func() IAggPlugin

type AggregateType string

// aggregate is the event sourcing equivalent of a record.
type aggregate struct {
	ID                core.IIdentity
	Version           int64
	AppliedEvents     []IEvt
	UncommittedEvents []IEvt
	Type              AggregateType
	withAppliedEvents bool
	state             model.IWriteModel
	executors         map[CommandType]ITryCmd
	appliers          map[EventType]IApplyEvt
}

// NewAggregate initializes a new empty Aggregate
func NewAggregate(at AggregateType, state model.IWriteModel) IAggregate {
	result := &aggregate{
		Version:           aggregateStartVersion,
		AppliedEvents:     make([]IEvt, 0, aggregateAppliedEventsInitialCap),
		UncommittedEvents: make([]IEvt, 0, aggregateUncommittedEventsInitialCap),
		withAppliedEvents: false,
		executors:         make(map[CommandType]ITryCmd, 0),
		appliers:          make(map[EventType]IApplyEvt, 0),
		Type:              at,
		state:             state,
	}
	return result
}

// SetID set aggregate ID
func (a *aggregate) SetID(identity core.IIdentity) IAggregate {
	a.ID = identity
	return a
}

// GetID get aggregate ID
func (a *aggregate) GetID() core.IIdentity {
	if a.ID == nil {
		panic(ErrTheAggregateHasNoID)
	}
	return a.ID
}

// GetAggregateType get aggregate AggregateType
func (a *aggregate) GetAggregateType() AggregateType {
	if a.Type == "" {
		panic(ErrInvalidAggregateType)
	}
	return a.Type
}

// GetVersion get aggregate version
func (a *aggregate) GetVersion() int64 {
	return a.Version
}

// ClearUncommittedEvents clear aggregate uncommitted Event's
func (a *aggregate) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]IEvt, 0, aggregateUncommittedEventsInitialCap)
}

// GetAppliedEvents get aggregate applied Event's
func (a *aggregate) GetAppliedEvents() []IEvt {
	return a.AppliedEvents
}

// SetAppliedEvents set aggregate applied Event's
func (a *aggregate) SetAppliedEvents(events []IEvt) {
	a.AppliedEvents = events
}

// GetUncommittedEvents get aggregate uncommitted Event's
func (a *aggregate) GetUncommittedEvents() []IEvt {
	return a.UncommittedEvents
}

// Load add existing events from event store to domain using IApplyEvt interface method
func (a *aggregate) Load(events []IEvt) error {
	for _, evt := range events {
		if err := a.ApplyEvent(evt, true); err != nil {
			return err
		}
		if a.withAppliedEvents {
			a.AppliedEvents = append(a.AppliedEvents, evt)
		}
		a.Version++
	}
	return nil
}

// Apply push event to domain uncommitted events using IApplyEvt method
func (a *aggregate) ApplyEvent(event IEvt, isCommitted bool) error {
	if event.GetAggregateId() == "" {
		return ErrEventHasNoAggregateID
	}
	if event.GetAggregateId() != a.GetID().Id() {
		return ErrInvalidAggregate
	}
	event.SetAggregateType(a.GetAggregateType())
	apply, err := a.getApplyEvt(event.GetEventType())
	if err != nil {
		return err
	}
	if err := apply.ApplyEvent(event, a.state); err != nil {
		return err
	}
	a.Version++
	event.SetVersion(a.GetVersion())
	if !isCommitted {
		a.UncommittedEvents = append(a.UncommittedEvents, event.(*Event))
	}
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, event)
	}
	return nil
}

// RaiseEvent push event to domain applied events using IApplyEvt method, used for load directly from eventstore
func (a *aggregate) RaiseEvent(event IEvt) error {
	if event.GetAggregateId() != a.GetID().Id() {
		return ErrInvalidAggregateID
	}
	if a.GetVersion() >= event.GetVersion() {
		return ErrInvalidEventVersion
	}
	event.SetAggregateType(a.GetAggregateType())
	if err := a.ApplyEvent(event, true); err != nil {
		return err
	}
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, event)
	}
	a.Version = event.GetVersion()
	return nil
}

// ToSnapshot prepare aggregate for saving Snapshot.
func (a *aggregate) ToSnapshot() {
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, a.UncommittedEvents...)
	}
	a.ClearUncommittedEvents()
}

func (a *aggregate) String() string {
	return fmt.Sprintf("Id: {%s}, Version: {%v}, Type: {%v}, AppliedEvents: {%v}, UncommittedEvents: {%v}",
		a.GetID(),
		a.GetVersion(),
		a.GetAggregateType(),
		len(a.GetAppliedEvents()),
		len(a.GetUncommittedEvents()))
}

func (a *aggregate) GetState() model.IWriteModel {
	return a.state
}

func (a *aggregate) TryCommand(ctx context.Context, cmd ICmd) (IEvt, dtos.IFbk) {
	fbk := dtos.NewFbk(cmd.GetAggregateID().Id(), -1, "")
	if cmd == nil {
		fbk.SetError(CommandCannotBeNil)
		return nil, fbk
	}
	if cmd.GetAggregateID() == nil {
		fbk.SetError(CommandMustHaveAggregateID)
		return nil, fbk
	}
	a.SetID(cmd.GetAggregateID())
	if cmd.GetCommandType() == "" {
		fbk.SetError(CommandTypeMustNotBeEmpty)
		return nil, fbk
	}
	exec, err := a.getExecCmd(cmd.GetCommandType())
	if err != nil {
		fbk.SetError(err.Error())
		return nil, fbk
	}
	e, f := exec.TryCommand(ctx, cmd)
	if !f.IsSuccess() {
		return nil, f
	}
	e.SetAggregateId(cmd.GetAggregateID().Id())
	er := a.ApplyEvent(e, false)
	if er != nil {
		f.SetError(er.Error())
	}
	// With payload
	f.SetJsonData(a.GetState())
	return e, f

}

func (a *aggregate) getApplyEvt(eventType EventType) (IApplyEvt, error) {
	applier := a.appliers[eventType]
	if applier == nil {
		return nil, fmt.Errorf(NoApplierForEvent, eventType)
	}
	return applier, nil
}

//getExecCmd returns the command executer for the specified CommandType.
func (a *aggregate) getExecCmd(commandType CommandType) (ITryCmd, error) {
	exec := a.executors[commandType]
	if exec == nil {
		return nil, fmt.Errorf(NoExecuterForCommand, commandType)
	}
	return exec, nil
}

//regTryCmd registers an TryCmd,
//we make sure nobody can inject at a taken slot, for safety
func (a *aggregate) regTryCmd(execute ITryCmd) {
	if execute == nil {
		return
	}
	existing, _ := a.getExecCmd(execute.GetCommandType())
	if existing == nil {
		execute.SetAggregate(a)
		a.executors[execute.GetCommandType()] = execute
	}
}

//regApplyEvt registers an event applier in the Aggregate
//we make sure nobody can inject at a taken slot, for safety
func (a *aggregate) regApplyEvt(apply IApplyEvt) {
	if apply == nil {
		return
	}
	existing, _ := a.getApplyEvt(apply.GetEventType())
	if existing == nil {
		apply.SetAggregate(a)
		a.appliers[apply.GetEventType()] = apply
	}
}

//Inject allows you to inject a series of aggregate ActorFtors
//and will inject each actor in the right map,
//depending on its type
func (a *aggregate) Inject(agg IAggregate, actors ...AggPluginFtor) IAggregate {
	for _, actFtor := range actors {
		act := actFtor()
		switch act.(type) {
		case ITryCmd:
			a.regTryCmd(act.(ITryCmd))
		case IApplyEvt:
			a.regApplyEvt(act.(IApplyEvt))
		default:
			continue
		}
	}
	return agg
}

// KnowsCmd returns true if
func (a *aggregate) KnowsCmd(topic CommandType) bool {
	return a.executors[topic] != nil
}
func (a *aggregate) KnowsEvt(topic EventType) bool {
	return a.appliers[topic] != nil
}

// IsAggregateFound checks the version. If it is != 0, the aggregate is found.
func IsAggregateFound(aggregate IAggregate) bool {
	return aggregate.GetVersion() != 0
}

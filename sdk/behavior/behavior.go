package behavior

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"sync"
)

const (
	startVersion                = -1 // used for EventStoreDB
	appliedEventsInitialCap     = 10
	uncommittedEventsInitialCap = 10
)

type ISetBehavior interface {
	SetBehavior(a IBehavior)
}

type IGetEvtType interface {
	GetEventType() EventType
}

type IGetCmdType interface {
	GetCommandType() CommandType
}

//IBehaviorPlugin  is an injector that allows us to inject ITryCmd and IApplyEvt injectors into the Aggregate in an elegant way.
type IBehaviorPlugin interface {
	ISetBehavior
}

// IApplyEvt is an IBehaviorPlugin injector that allows us to inject Event Appliers into the Aggregate
type IApplyEvt interface {
	IBehaviorPlugin
	IGetEvtType
	ApplyEvent(state schema.ISchema, evt IEvt) error
}

// ITryCmd is an IBehaviorPlugin injector that allows us to inject Command Executors into the Aggregate
type ITryCmd interface {
	IBehaviorPlugin
	IGetCmdType
	TryCommand(ctx context.Context, command ICmd) (IEvt, contract.IFbk)
}

// IBehavior is the injector for Aggregates
// In an Event Sourced application, the Aggregate can be considered as the equivalent of the ActiveRecord in classic CrUD (Create, Upsert, Delete) applications.
// The Aggregate can be considered as the heart of an ES system, that unites State with Behavior.
// Its main responsibilities are:
//  1. to build a (volatile) State from an ordered list of previously committed Events, that are sourced from the Event Stream that is identified by the Id.
//  2. CanAcceptName Command requests and applying business logic that checks whether the Command (ICmd) is allowed to be executed or not,
//     according to a number of Specifications, the Current State (see 1.) and the Command's FactPayload
//  3. If Command execution is allowed, Raise a new Event and ApplyEvent it to itself, as to update the Current State to the New State.
// Specifically for GO-SCREAM CMD Applications, given their modular nature, we rely on Aggregate Composition.
// Aggregate Composition is a technique that allows us to inject
// a series of BehaviorPluginFtor functors that create a feature's IApplyEvt and ITryCmd injectors,
// in order to compose an Aggregate that has all the capabilities required to process the Event Stream.
type IBehavior interface {
	ISimpleBehavior
	Inject(agg IBehavior, actors ...BehaviorPluginFtor) IBehavior
	KnowsCmd(topic CommandType) bool
	KnowsEvt(topic EventType) bool
}

// ISimpleBehavior is an injector that abstracts the implementation of a traditional monolithic Aggregate.
type ISimpleBehavior interface {
	IGetBehaviorType
	String() string
	TryCommand(ctx context.Context, command ICmd) (IEvt, contract.IFbk)
	ApplyEvent(event IEvt, isCommitted bool) error
	GetState() schema.IModel
	GetUncommittedEvents() []IEvt
	ClearUncommittedEvents()
	SetAppliedEvents(events []IEvt)
	GetAppliedEvents() []IEvt
	ToSnapshot()
	RaiseEvent(event IEvt) error
	GetVersion() int64
	GetID() schema.IIdentity
	SetID(identity schema.IIdentity) IBehavior
}

type BehaviorFtor func() IBehavior
type GenBehaviorFtor[TDoc schema.ISchema] func() IBehavior
type BehaviorBuilder func() IBehavior
type BehaviorPluginFtor func() IBehaviorPlugin

type BehaviorType string

// behavior is the event sourcing equivalent of a record.
type behavior struct {
	ID                schema.IIdentity
	Version           int64
	AppliedEvents     []IEvt
	UncommittedEvents []IEvt
	Type              BehaviorType
	withAppliedEvents bool
	state             schema.IModel
	executors         map[CommandType]ITryCmd
	appliers          map[EventType]IApplyEvt
}

// NewBehavior initializes a new empty Aggregate
func NewBehavior(behaviorType BehaviorType, state schema.IModel) IBehavior {
	result := &behavior{
		Version:           startVersion,
		AppliedEvents:     make([]IEvt, 0, appliedEventsInitialCap),
		UncommittedEvents: make([]IEvt, 0, uncommittedEventsInitialCap),
		withAppliedEvents: false,
		executors:         make(map[CommandType]ITryCmd, 0),
		appliers:          make(map[EventType]IApplyEvt, 0),
		Type:              behaviorType,
		state:             state,
	}
	return result
}

// SetID set behavior ID
func (a *behavior) SetID(identity schema.IIdentity) IBehavior {
	a.ID = identity
	return a
}

// GetID get behavior ID
func (a *behavior) GetID() schema.IIdentity {
	if a.ID == nil {
		panic(ErrTheBehaviorHasNoID)
	}
	return a.ID
}

// GetAggregateType get behavior BehaviorType
func (a *behavior) GetBehaviorType() BehaviorType {
	if a.Type == "" {
		panic(ErrInvalidBehaviorType)
	}
	return a.Type
}

// GetVersion get behavior version
func (a *behavior) GetVersion() int64 {
	return a.Version
}

// ClearUncommittedEvents clear behavior uncommitted Event's
func (a *behavior) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]IEvt, 0, uncommittedEventsInitialCap)
}

// GetAppliedEvents get behavior applied Event's
func (a *behavior) GetAppliedEvents() []IEvt {
	return a.AppliedEvents
}

// SetAppliedEvents set behavior applied Event's
func (a *behavior) SetAppliedEvents(events []IEvt) {
	a.AppliedEvents = events
}

// GetUncommittedEvents get behavior uncommitted Event's
func (a *behavior) GetUncommittedEvents() []IEvt {
	return a.UncommittedEvents
}

// Load add existing events from event store to domain using IApplyEvt interface method
func (a *behavior) Load(events []IEvt) error {
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

var aMutex = &sync.Mutex{}

// ApplyEvent push event to domain uncommitted events using IApplyEvt method
func (a *behavior) ApplyEvent(event IEvt, isCommitted bool) error {
	aMutex.Lock()
	defer aMutex.Unlock()
	if event.GetBehaviorId() == "" {
		return ErrEventHasNoBehaviorID
	}
	if event.GetBehaviorId() != a.GetID().Id() {
		return ErrInvalidBehavior
	}
	event.SetBehaviorType(a.GetBehaviorType())
	apply, err := a.getApplyEvt(event.GetEventType())
	if err != nil {
		return err
	}
	if err := apply.ApplyEvent(a.state, event); err != nil {
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

var raiseMutex = sync.Mutex{}

// RaiseEvent push event to domain applied events using IApplyEvt method, used for load directly from eventstore
func (a *behavior) RaiseEvent(event IEvt) error {
	raiseMutex.Lock()
	defer raiseMutex.Unlock()
	if event.GetBehaviorId() != a.GetID().Id() {
		return ErrInvalidBehaviorID
	}
	if a.GetVersion() >= event.GetVersion() {
		return ErrInvalidEventVersion
	}
	event.SetBehaviorType(a.GetBehaviorType())
	if err := a.ApplyEvent(event, true); err != nil {
		return err
	}
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, event)
	}
	a.Version = event.GetVersion()
	return nil
}

// ToSnapshot prepare behavior for saving Snapshot.
func (a *behavior) ToSnapshot() {
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, a.UncommittedEvents...)
	}
	a.ClearUncommittedEvents()
}

func (a *behavior) String() string {
	return fmt.Sprintf("Id: {%s}, Version: {%v}, Type: {%v}, AppliedEvents: {%v}, UncommittedEvents: {%v}",
		a.GetID(),
		a.GetVersion(),
		a.GetBehaviorType(),
		len(a.GetAppliedEvents()),
		len(a.GetUncommittedEvents()))
}

func (a *behavior) GetState() schema.IModel {
	return a.state
}

var tryMutex = &sync.Mutex{}

func (a *behavior) TryCommand(ctx context.Context, cmd ICmd) (IEvt, contract.IFbk) {
	tryMutex.Lock()
	defer tryMutex.Unlock()
	fbk := contract.NewFbk(cmd.GetBehaviorID().Id(), -1, "")
	if cmd == nil {
		fbk.SetError(CommandCannotBeNil)
		return nil, fbk
	}
	if cmd.GetBehaviorID() == nil {
		fbk.SetError(CommandMustHaveBehaviorID)
		return nil, fbk
	}
	a.SetID(cmd.GetBehaviorID())
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
	e.SetBehaviorId(cmd.GetBehaviorID().Id())
	er := a.ApplyEvent(e, false)
	if er != nil {
		f.SetError(er.Error())
	}
	// With payload
	f.SetStatus(a.GetState().GetStatus())
	f.SetPayload(a.GetState())
	return e, f

}

func (a *behavior) getApplyEvt(eventType EventType) (IApplyEvt, error) {
	applier := a.appliers[eventType]
	if applier == nil {
		return nil, fmt.Errorf(NoApplierForEvent, eventType)
	}
	return applier, nil
}

//getExecCmd returns the command executer for the specified CommandType.
func (a *behavior) getExecCmd(commandType CommandType) (ITryCmd, error) {
	exec := a.executors[commandType]
	if exec == nil {
		return nil, fmt.Errorf(NoExecuterForCommand, commandType)
	}
	return exec, nil
}

//regTryCmd registers an TryCmd,
//we make sure nobody can inject at a taken slot, for safety
func (a *behavior) regTryCmd(execute ITryCmd) {
	if execute == nil {
		return
	}
	existing, _ := a.getExecCmd(execute.GetCommandType())
	if existing == nil {
		execute.SetBehavior(a)
		a.executors[execute.GetCommandType()] = execute
	}
}

//regApplyEvt registers an event applier in the Aggregate
//we make sure nobody can inject at a taken slot, for safety
func (a *behavior) regApplyEvt(apply IApplyEvt) {
	if apply == nil {
		return
	}
	existing, _ := a.getApplyEvt(apply.GetEventType())
	if existing == nil {
		apply.SetBehavior(a)
		a.appliers[apply.GetEventType()] = apply
	}
}

//Inject allows you to inject a series of behavior ActorFtors
//and will inject each actor in the right map,
//depending on its type
func (a *behavior) Inject(agg IBehavior, actors ...BehaviorPluginFtor) IBehavior {
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
func (a *behavior) KnowsCmd(topic CommandType) bool {
	return a.executors[topic] != nil
}
func (a *behavior) KnowsEvt(topic EventType) bool {
	return a.appliers[topic] != nil
}

// IsBehaviorFound checks the version. If it is != 0, the behavior is found.
func IsBehaviorFound(behavior IBehavior) bool {
	return behavior.GetVersion() != 0
}

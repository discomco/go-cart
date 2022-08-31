package domain

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/model"
)

const (
	AggregateCannotBeNil = "domain cannot be nil"
)

var (
	ErrAggregateCannotBeNil = fmt.Errorf(AggregateCannotBeNil)
)

type applyEvtToStateFunc func(evt IEvt, state model.IWriteModel) error

type ApplyEvt struct {
	aggregate       IAggregate
	eventType       EventType
	applyEvtToState applyEvtToStateFunc
}

// TODO: Check if we can ommit this construction
func (a *ApplyEvt) ApplyEvent(event IEvt, state model.IWriteModel) error {
	return a.applyEvtToState(event, state)
}

func (a *ApplyEvt) GetAggregate() IAggregate {
	return a.aggregate
}

func (a *ApplyEvt) GetEventType() EventType {
	return a.eventType
}

func (a *ApplyEvt) SetAggregate(agg IAggregate) {
	a.aggregate = agg
}

// NewApplyEvt lets you create an Event Applier and requires that you pass an applyEvtToStateFunc function.
// your Event Applier is automatically injected into the Aggregate.
func NewApplyEvt(
	eventType EventType,
	applyEvtToState applyEvtToStateFunc,
) *ApplyEvt {
	result := &ApplyEvt{
		eventType:       eventType,
		applyEvtToState: applyEvtToState,
	}
	return result
}

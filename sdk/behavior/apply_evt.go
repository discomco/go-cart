package behavior

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/schema"
)

const (
	CannotBeNil = "behavior cannot be nil"
)

var (
	ErrBehaviorCannotBeNil = fmt.Errorf(CannotBeNil)
)

type applyEvtToStateFunc func(evt IEvt, state schema.IWriteSchema) error

type ApplyEvt struct {
	behavior        IBehavior
	eventType       EventType
	applyEvtToState applyEvtToStateFunc
}

func (a *ApplyEvt) ApplyEvent(event IEvt, state schema.IWriteSchema) error {
	return a.applyEvtToState(event, state)
}

func (a *ApplyEvt) GetAggregate() IBehavior {
	return a.behavior
}

func (a *ApplyEvt) GetEventType() EventType {
	return a.eventType
}

func (a *ApplyEvt) SetAggregate(agg IBehavior) {
	a.behavior = agg
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

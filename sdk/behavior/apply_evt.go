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

type FApply func(evt IEvt, state schema.ISchema) error

type ApplyEvt struct {
	behavior  IBehavior
	eventType EventType
	fApply    FApply
}

func (a *ApplyEvt) ApplyEvent(event IEvt, state schema.ISchema) error {
	return a.fApply(event, state)
}

func (a *ApplyEvt) GetAggregate() IBehavior {
	return a.behavior
}

func (a *ApplyEvt) GetEventType() EventType {
	return a.eventType
}

func (a *ApplyEvt) SetBehavior(agg IBehavior) {
	a.behavior = agg
}

// NewApplyEvt lets you create an Event Applier and requires that you pass an FApply function.
// your Event Applier is automatically injected into the Aggregate.
func NewApplyEvt(
	eventType EventType,
	fApply FApply,
) *ApplyEvt {
	result := &ApplyEvt{
		eventType: eventType,
		fApply:    fApply,
	}
	return result
}

package behavior

import (
	"fmt"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/pkg/errors"
)

func EvtToList() behavior.FEvt2Schema[IEvt, schema.QuadraticList] {
	return func(evt IEvt, list *schema.QuadraticList) error {
		return evtToList(evt, list)
	}
}

func evtToList(evt IEvt, list *schema.QuadraticList) error {
	calcId := evt.GetBehaviorId()
	calcIt := list.GetItem(calcId)
	var pl contract.Payload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(initialize_calc.evtToList) failed to get payload from event %v", evt)
	}
	calcIt.Status = doc.Initialized
	calcIt.Equation = fmt.Sprintf("(%.2f)x^2+(%.2f)x+(%.2f)=0", pl.Input.A, pl.Input.B, pl.Input.C)
	list.Items[calcId] = calcIt
	return nil
}

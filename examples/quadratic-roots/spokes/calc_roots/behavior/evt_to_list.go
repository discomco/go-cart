package behavior

import (
	"fmt"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-status"
	"github.com/pkg/errors"
)

func EvtToList() behavior.Evt2DocFunc[IEvt, schema.QuadraticList] {
	return func(evt IEvt, list *schema.QuadraticList) error {
		return evtToList(evt, list)
	}
}

func evtToList(evt IEvt, list *schema.QuadraticList) error {
	calcId := evt.GetBehaviorId()
	calcIt := list.GetItem(calcId)
	var pl contract.FactPayload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(calc_roots.evtToList) failed to get payload from event %v", evt)
	}
	status.SetStatus(&calcIt.Status, doc.RootsCalculated)
	calcIt.Discriminator = fmt.Sprintf("%.2f", pl.Output.D)
	if pl.Output.D < 0 {
		calcIt.X1 = fmt.Sprintf("x1=%.2f+%.2fj", pl.Output.X1.Real, pl.Output.X1.Imaginary)
		calcIt.X2 = fmt.Sprintf("x1=%.2f+%.2fj", pl.Output.X2.Real, pl.Output.X2.Imaginary)
	} else {
		calcIt.X1 = fmt.Sprintf("x1=%.2f", pl.Output.X1.Real)
		calcIt.X2 = fmt.Sprintf("x2=%.2f", pl.Output.X2.Real)
	}
	list.Items[calcId] = calcIt
	return nil
}

package behavior

import (
	app_schema "github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-status"
	"github.com/pkg/errors"
)

type IApplyEvt interface {
	behavior.IApplyEvt
}

type apply struct {
	*behavior.ApplyEvt
}

func (a *apply) fApply(state schema.ISchema, evt behavior.IEvt) error {
	// EXTRACT FactPayload
	var pl contract.FactPayload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(applyEvent) could not extract payload")
	}
	s := state.(*app_schema.QuadraticDoc)
	ID, _ := evt.GetBehaviorID()
	s.ID = ID.(*schema.Identity)
	s.Output = pl.Output
	status.SetStatus(&s.Status, doc.RootsCalculated)
	return err
}

func newApply() IApplyEvt {
	a := &apply{}
	b := behavior.NewApplyEvt(EvtTopic, a.fApply)
	a.ApplyEvt = b
	return a
}

func ApplyEvt() behavior.IBehaviorPlugin {
	return newApply()
}

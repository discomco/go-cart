package behavior

import (
	read_model "github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

type IApplyEvt interface {
	behavior.IApplyEvt
}

type apply struct {
	*behavior.ApplyEvt
}

func (a *apply) applyEvt(evt behavior.IEvt, state schema.IModel) error {
	// EXTRACT Payload
	var pl contract.Payload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(applyEvent) could not extract payload")
	}
	s := state.(*read_model.GameDoc)
	if pl.Details != nil {
		s.Details = pl.Details
	}
	return err
}

func newApply() IApplyEvt {
	a := &apply{}
	b := behavior.NewApplyEvt(EVT_TOPIC, a.applyEvt)
	a.ApplyEvt = b
	return a
}

func ApplyEvt() behavior.IBehaviorPlugin {
	return newApply()
}

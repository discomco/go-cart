package behavior

import (
	read_model "github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_settings/contract"
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

func (a *apply) fApply(state schema.ISchema, evt behavior.IEvt) error {
	// EXTRACT Payload
	var pl contract.Payload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(applyEvent) could not extract payload")
	}
	s := state.(*read_model.GameDoc)
	if pl.Settings != nil {
		s.Settings = pl.Settings
	}
	return err
}

func newApply() IApplyEvt {
	a := &apply{}
	b := behavior.NewFapply(EVT_TOPIC, a.fApply)
	a.ApplyEvt = b
	return a
}

func ApplyEvt() behavior.IBehaviorPlugin {
	return newApply()
}

package behavior

import (
	read_model "github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
	go_status "github.com/discomco/go-status"
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
	ID, _ := evt.GetBehaviorID()
	s.ID = ID.(*schema.Identity)
	s.Details = pl.Details
	go_status.SetStatus(&s.Status, doc.Initialized)
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

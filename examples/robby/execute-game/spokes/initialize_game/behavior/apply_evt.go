package behavior

import (
	read_model "github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

type IApplyEvt interface {
	behavior.IApplyEvt
}

type apply struct {
	*behavior.ApplyEvt
}

func (a *apply) applyEvt(evt behavior.IEvt, state schema.IWriteSchema) error {
	// EXTRACT Payload
	var pl contract.Payload
	err := evt.GetPayload(&pl)
	if err != nil {
		return errors.Wrapf(err, "(applyEvent) could not extract payload")
	}
	s := state.(*read_model.GameDoc)
	ID, _ := evt.GetAggregateID()
	s.ID = ID.(*schema.Identity)
	s.Details = pl.Details
	status.SetFlag(&s.Status, doc.Initialized)
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

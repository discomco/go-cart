package behavior

import (
	read_model "github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/model"
	"github.com/pkg/errors"
)

type IApplyEvt interface {
	domain.IApplyEvt
}

type apply struct {
	*domain.ApplyEvt
}

func (a *apply) applyEvt(evt domain.IEvt, state model.IWriteModel) error {
	// EXTRACT Payload
	var pl contract.Payload
	err := evt.GetJsonData(&pl)
	if err != nil {
		return errors.Wrapf(err, "(applyEvent) could not extract payload")
	}
	s := state.(*read_model.GameDoc)
	ID, _ := evt.GetAggregateID()
	s.ID = ID.(*core.Identity)
	s.Details = pl.Details
	status.SetFlag(&s.Status, doc.Initialized)
	return err
}

func newApply() IApplyEvt {
	a := &apply{}
	b := domain.NewApplyEvt(EVT_TOPIC, a.applyEvt)
	a.ApplyEvt = b
	return a
}

func ApplyEvt() domain.IAggPlugin {
	return newApply()
}

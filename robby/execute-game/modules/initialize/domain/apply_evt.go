package domain

import (
	model2 "github.com/discomco/go-cart/robby/execute-game/-shared/model"
	"github.com/discomco/go-cart/robby/execute-game/-shared/model/root"
	"github.com/discomco/go-cart/robby/execute-game/modules/initialize/dtos"
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
	var pl dtos.Payload
	err := evt.SetJsonData(&pl)
	if err != nil {
		return errors.Wrapf(err, "(applyEvent) could not extract payload")
	}

	s := state.(*model2.Root)
	ID, _ := evt.GetAggregateID()
	s.ID = ID.(*core.Identity)
	s.Details = pl.Details
	status.SetFlag(&s.Status, root.Initialized)
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

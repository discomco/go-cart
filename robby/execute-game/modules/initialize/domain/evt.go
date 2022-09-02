package domain

import (
	"github.com/discomco/go-cart/robby/execute-game/modules/initialize/dtos"
	"github.com/discomco/go-cart/sdk/domain"
)

type IEvt interface {
	domain.IEvt
}

func NewEvt(aggregate domain.IAggregate, payload dtos.Payload) IEvt {
	e := domain.NewEvt(aggregate, EVT_TOPIC)
	e.SetJsonData(payload)
	return e
}

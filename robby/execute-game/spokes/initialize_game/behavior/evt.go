package behavior

import (
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/domain"
)

type IEvt interface {
	domain.IEvt
}

func NewEvt(aggregate domain.IAggregate, payload contract.Payload) IEvt {
	e := domain.NewEvt(aggregate, EVT_TOPIC)
	e.SetJsonData(payload)
	return e
}

package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/behavior"
)

type IEvt interface {
	behavior.IEvt
}

func NewEvt(aggregate behavior.IBehavior, payload contract.Payload) IEvt {
	e := behavior.NewEvt(aggregate, EvtTopic)
	e.SetPayload(payload)
	return e
}

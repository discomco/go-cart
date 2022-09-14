package behavior

import (
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/behavior"
)

type IEvt interface {
	behavior.IEvt
}

func NewEvt(aggregate behavior.IBehavior, payload contract.Payload) IEvt {
	e := behavior.NewEvt(aggregate, EVT_TOPIC)
	e.SetPayload(payload)
	return e
}

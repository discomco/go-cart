package comps

import (
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IRequester interface {
	nats.IRequester[contract.IHope]
}

func Requester() (IRequester, error) {
	return nats.NewRequester[contract.IHope](contract.HOPE_TOPIC)
}

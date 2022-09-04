package reactors

import (
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IRequester interface {
	nats.INATSRequester[contract.IHope]
}

func Requester() (IRequester, error) {
	return nats.NewRequester[contract.IHope](contract.HOPE_TOPIC)
}

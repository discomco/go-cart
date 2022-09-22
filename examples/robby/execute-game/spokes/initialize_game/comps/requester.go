package comps

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IRequester interface {
	nats.IRequester[contract.IHope]
}

func Requester() (IRequester, error) {
	return nats.NewRequester[contract.IHope](contract.HopeTopic)
}

func NewRequester() comps.RequesterFtor {
	return func() (comps.IRequester, error) {
		return Requester()
	}
}

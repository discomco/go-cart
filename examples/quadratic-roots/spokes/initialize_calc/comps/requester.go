package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IRequester interface {
	nats.IRequester[contract.IHope]
}

func Requester() (IRequester, error) {
	return nats.NewRequester[contract.IHope](contract.HopeTopic)
}

func RequesterFtor() comps.GenRequesterFtor[contract.IHope] {
	return func() (comps.IGenRequester[contract.IHope], error) {
		return Requester()
	}
}

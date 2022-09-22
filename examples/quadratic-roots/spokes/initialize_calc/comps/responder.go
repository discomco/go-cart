package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	behavior2 "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IResponder interface {
	nats.IResponder[contract.IHope, behavior.ICmd]
}

func Responder(h2c behavior2.Hope2CmdFunc[contract.IHope, behavior.ICmd]) (IResponder, error) {
	return nats.NewResponder[contract.IHope, behavior.ICmd](contract.HopeTopic, h2c)
}

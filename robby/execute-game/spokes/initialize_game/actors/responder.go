package actors

import (
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IResponder interface {
	nats.INATSResponder[contract.IHope, domain.ICmd]
}

func Responder(hope2Cmd domain.Hope2CmdFunc[contract.IHope, behavior.ICmd]) (IResponder, error) {
	return nats.NewResponder[contract.IHope, behavior.ICmd](
		contract.HOPE_TOPIC,
		hope2Cmd,
	)
}

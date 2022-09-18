package comps

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/behavior"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/drivers/nats"
)

type IResponder interface {
	nats.IResponder[contract.IHope, behavior.ICmd]
}

func Responder(hope2Cmd sdk_behavior.Hope2CmdFunc[contract.IHope, behavior.ICmd]) (IResponder, error) {
	return nats.NewResponder[contract.IHope, behavior.ICmd](
		contract.HOPE_TOPIC,
		hope2Cmd,
	)
}

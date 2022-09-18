package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	sdk_contract "github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

func Hope2Cmd() behavior.Hope2CmdFunc[contract.IHope, ICmd] {
	return func(hope *sdk_contract.Dto) (ICmd, error) {
		var pl contract.Payload
		err := hope.GetPayload(&pl)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to extract payload from hope: %v", err)
		}
		gameID, err := schema.IdentityFromPrefixedId(hope.GetId())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to extract gameId from hope: %v", err)
		}
		cmd, err := NewCmd(gameID, pl)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create Cmd from Hope: %v", err)
		}
		return cmd, nil
	}
}

package behavior

import (
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/pkg/errors"
)

func Hope2Cmd() domain.Hope2CmdFunc[contract.IHope, ICmd] {
	return func(hope *dtos.Dto) (ICmd, error) {
		var pl contract.Payload
		err := hope.GetJsonData(&pl)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to extract payload from hope: %v", err)
		}
		gameID, err := core.IdentityFromPrefixedId(pl.GameId)
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

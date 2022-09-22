package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

type ICmd interface {
	behavior.ICmd
}

func NewCmd(aggID schema.IIdentity, payload contract.Payload) (ICmd, error) {
	data, err := convert.Any2Data(payload)
	if err != nil {
		return nil, errors.Wrapf(err, "(NewCmd) failed to convert payload %v", payload)
	}
	c, err := behavior.NewCmd(aggID, CmdTopic, data)
	if err != nil {
		return nil, errors.Wrapf(err, "(NewCmd) failed to create ICmd for topic [%v]", CmdTopic)
	}
	return c, err
}

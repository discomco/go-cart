package behavior

import (
	contract "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

type ICmd interface {
	behavior.ICmd
}

func NewCmd(calcID schema.IIdentity, payload contract.HopePayload) (ICmd, error) {
	data, err := convert.Any2Data(payload)
	if err != nil {
		return nil, errors.Wrapf(err, "(NewCmd) failed to convert payload %v", payload)
	}
	c, err := behavior.NewCmd(calcID, CmdTopic, data)
	if err != nil {
		return nil, errors.Wrapf(err, "(NewCmd) failed to create ICmd for topic [%v]", CmdTopic)
	}
	return c, err
}

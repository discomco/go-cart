package domain

import (
	"github.com/discomco/go-cart/robby/execute-game/modules/initialize/dtos"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/pkg/errors"
)

type ICmd interface {
	domain.ICmd
}

func NewCmd(aggID core.IIdentity, payload dtos.Payload) (ICmd, error) {
	data, err := convert.Any2Data(payload)
	if err != nil {
		return nil, errors.Wrapf(err, "(NewCmd) failed to convert payload %v", payload)
	}
	c, err := domain.NewCmd(aggID, CMD_TOPIC, data)
	if err != nil {
		return nil, errors.Wrapf(err, "(NewCmd) failed to create ICmd for topic [%v]", CMD_TOPIC)
	}
	return c, err
}

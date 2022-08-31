package cmd_must

import (
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
)

func HaveAggregateID(cmd domain.ICmd, fbk dtos.IFbk) {
	if cmd.GetAggregateID() == nil {
		fbk.SetError(domain.CommandMustHaveAggregateID)
	}
}

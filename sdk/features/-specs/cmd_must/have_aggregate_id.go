package cmd_must

import (
	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/dtos"
)

func HaveAggregateID(cmd domain.ICmd, fbk dtos.IFbk) {
	if cmd.GetAggregateID() == nil {
		fbk.SetError(domain.CommandMustHaveAggregateID)
	}
}

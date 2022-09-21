package cmd_must

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/contract"
)

func HaveAggregateID(cmd behavior.ICmd, fbk contract.IFbk) {
	if cmd.GetBehaviorID() == nil {
		fbk.SetError(behavior.CommandMustHaveBehaviorID)
	}
}

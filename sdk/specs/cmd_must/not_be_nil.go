package cmd_must

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/contract"
)

func NotBeNil(cmd behavior.ICmd, fbk contract.IFbk) {
	if cmd == nil {
		fbk.SetError(behavior.CommandCannotBeNil)
	}
}

package cmd_must

import (
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
)

func NotBeNil(cmd domain.ICmd, fbk dtos.IFbk) {
	if cmd == nil {
		fbk.SetError(domain.CommandCannotBeNil)
	}
}

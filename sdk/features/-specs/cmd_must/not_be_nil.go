package cmd_must

import (
	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/dtos"
)

func NotBeNil(cmd domain.ICmd, fbk dtos.IFbk) {
	if cmd == nil {
		fbk.SetError(domain.CommandCannotBeNil)
	}
}

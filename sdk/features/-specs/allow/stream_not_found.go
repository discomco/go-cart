package allow

import (
	sdk_errors "github.com/discomco/go-cart/core/errors"
	"github.com/discomco/go-cart/dtos"
	"github.com/pkg/errors"
)

func StreamNotFound(err error, fbk dtos.IFbk) {
	if err != nil && !errors.Is(err, sdk_errors.ErrStreamNotFound) {
		fbk.SetError(err.Error())
	}
}
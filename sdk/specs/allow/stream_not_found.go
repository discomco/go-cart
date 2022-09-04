package allow

import (
	"github.com/discomco/go-cart/sdk/contract"
	sdk_errors "github.com/discomco/go-cart/sdk/core/errors"
	"github.com/pkg/errors"
)

func StreamNotFound(err error, fbk contract.IFbk) {
	if err != nil && !errors.Is(err, sdk_errors.ErrStreamNotFound) {
		fbk.SetError(err.Error())
	}
}

package state_must

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema"
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema/root"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	"github.com/discomco/go-cart/sdk/dtos"
	sdk_model "github.com/discomco/go-cart/sdk/model"
)

const AllReadyInitialized = "allReadyInitialized"

func NotBeInitialized(s sdk_model.IWriteModel, fbk dtos.IFbk) {
	state := s.(*schema.Root)
	if status.HasFlag(state.Status, root.Initialized) {
		fbk.SetError(AllReadyInitialized)
	}
}

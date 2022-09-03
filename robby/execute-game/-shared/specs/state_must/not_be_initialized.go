package state_must

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"
	"github.com/discomco/go-cart/robby/execute-game/-shared/model/root"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	"github.com/discomco/go-cart/sdk/dtos"
	sdk_model "github.com/discomco/go-cart/sdk/model"
)

const AllReadyInitialized = "allReadyInitialized"

func NotBeInitialized(s sdk_model.IWriteModel, fbk dtos.IFbk) {
	state := s.(*model.Root)
	if status.HasFlag(state.Status, root.Initialized) {
		fbk.SetError(AllReadyInitialized)
	}
}

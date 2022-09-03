package state_must

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	"github.com/discomco/go-cart/sdk/dtos"
	sdk_model "github.com/discomco/go-cart/sdk/model"
)

const AllReadyInitialized = "allReadyInitialized"

func NotBeInitialized(s sdk_model.IWriteModel, fbk dtos.IFbk) {
	state := s.(*schema.GameDoc)
	if status.HasFlag(state.Status, doc.Initialized) {
		fbk.SetError(AllReadyInitialized)
	}
}

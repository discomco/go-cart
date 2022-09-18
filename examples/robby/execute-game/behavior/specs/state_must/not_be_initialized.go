package state_must

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/utils/status"
	sdk_model "github.com/discomco/go-cart/sdk/schema"
)

const AllReadyInitialized = "allReadyInitialized"

func NotBeInitialized(s sdk_model.IWriteSchema, fbk contract.IFbk) {
	state := s.(*schema.GameDoc)
	if status.HasFlag(state.Status, doc.Initialized) {
		fbk.SetError(AllReadyInitialized)
	}
}

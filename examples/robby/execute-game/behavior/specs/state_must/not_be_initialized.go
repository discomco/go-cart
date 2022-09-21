package state_must

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/contract"
	sdk_model "github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-status"
)

const AllReadyInitialized = "allReadyInitialized"

func NotBeInitialized(s sdk_model.IModel, fbk contract.IFbk) {
	state := s.(*schema.GameDoc)
	if status.HasStatus(state.Status, doc.Initialized) {
		fbk.SetError(AllReadyInitialized)
	}
}

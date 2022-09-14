package state_must

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/utils/status"
)

const NotYetInitialized = "not yet initialized"

func BeInitialized(state *schema.GameDoc, fbk contract.IFbk) {
	if status.NotHasFlag(state.Status, doc.Initialized) {
		fbk.SetError(NotYetInitialized)
	}
}

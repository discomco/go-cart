package state_must

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-status"
)

const NotYetInitialized = "not yet initialized"

func BeInitialized(state *schema.GameDoc, fbk contract.IFbk) {
	if status.NotHasStatus(state.Status, doc.Initialized) {
		fbk.SetError(NotYetInitialized)
	}
}

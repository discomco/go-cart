package behavior

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	"github.com/discomco/go-cart/sdk/behavior"
)

type IEvt interface {
	behavior.IEvt
}

// NewEvt creates a new instance of IEvt
func NewEvt(beh behavior.IBehavior, payload contract.FactPayload) IEvt {
	e := behavior.NewEvt(beh, EvtTopic)
	e.SetPayload(payload)
	return e
}

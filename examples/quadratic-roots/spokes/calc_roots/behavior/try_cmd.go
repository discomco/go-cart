package behavior

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/specs/doc_must"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	calc_roots "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/contract"
)

type ITryCmd interface {
	behavior.ITryCmd
}

func TryCmd() behavior.IBehaviorPlugin {
	return newTry()
}

type try struct {
	*behavior.TryCmd
}

func (t *try) fRaise(ctx context.Context, cmd behavior.ICmd) (behavior.IEvt, contract.IFbk) {
	// Initializations
	behID := cmd.GetBehaviorID()
	fbk := contract.NewFbk(behID.Id(), -1, "")
	agg := t.GetBehavior()
	state := agg.GetState()

	// SPECIFICATIONS
	doc_must.BeInitialized(state, fbk)
	if !fbk.IsSuccess() {
		return nil, fbk
	}

	// PREPARE EVENT
	output := calcRoots(state.(*schema.QuadraticDoc).Input)
	pl := calc_roots.NewFactPayload(output)
	evt := NewEvt(agg, *pl)

	// RAISE Event
	return evt, fbk
}

func newTry() *try {
	t := &try{}
	b := behavior.NewTryCmd(CmdTopic, t.fRaise)
	t.TryCmd = b
	return t
}

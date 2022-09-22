package behavior

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/specs/doc_must"
	initialize_calc "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/pkg/errors"
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
	doc_must.NotBeInitialized(state, fbk)
	if !fbk.IsSuccess() {
		return nil, fbk
	}

	// PREPARE EVENT
	var pl initialize_calc.Payload
	err := cmd.GetJsonPayload(&pl)
	if err != nil {
		e := fmt.Sprint(errors.Wrapf(err, "(initialize_calc.fRaise) could not extract payload"))
		fbk.SetError(e)
	}
	evt := NewEvt(agg, pl)

	// RAISE Event
	return evt, fbk
}

func newTry() *try {
	t := &try{}
	b := behavior.NewTryCmd(CmdTopic, t.fRaise)
	t.TryCmd = b
	return t
}

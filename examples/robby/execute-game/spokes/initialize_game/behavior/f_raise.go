package behavior

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/examples/robby/execute-game/behavior/specs/state_must"
	initialize_game "github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
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
	aggID := cmd.GetBehaviorID()
	fbk := contract.NewFbk(aggID.Id(), -1, "")
	agg := t.GetAggregate()
	state := agg.GetState()
	// SPECIFICATIONS
	state_must.NotBeInitialized(state, fbk)
	if !fbk.IsSuccess() {
		return nil, fbk
	}

	// PREPARE EVENT
	var pl initialize_game.Payload
	err := cmd.GetJsonPayload(&pl)
	if err != nil {
		e := fmt.Sprint(errors.Wrapf(err, "(initialize.fRaise) could not extract payload"))
		fbk.SetError(e)
	}
	evt := NewEvt(agg, pl)
	// RAISE Event
	return evt, fbk
}

func newTry() *try {
	t := &try{}
	b := behavior.NewTryCmd(CMD_TOPIC, t.fRaise)
	t.TryCmd = b
	return t
}

package behavior

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/robby/execute-game/behavior/specs/state_must"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	change_game_settings "github.com/discomco/go-cart/robby/execute-game/spokes/change_game_settings/contract"
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

func (t *try) raiseEvent(ctx context.Context, cmd behavior.ICmd) (behavior.IEvt, contract.IFbk) {
	// Initializations
	aggID := cmd.GetAggregateID()
	fbk := contract.NewFbk(aggID.Id(), -1, "")
	agg := t.GetAggregate()
	state := agg.GetState()
	// SPECIFICATIONS
	state_must.BeInitialized(state.(*schema.GameDoc), fbk)
	if !fbk.IsSuccess() {
		return nil, fbk
	}

	// PREPARE EVENT
	var pl change_game_settings.Payload
	err := cmd.GetJsonPayload(&pl)
	if err != nil {
		e := fmt.Sprint(errors.Wrapf(err, "(changeEventDetails.raiseEvent) could not extract payload"))
		fbk.SetError(e)
	}
	evt := NewEvt(agg, pl)
	// RAISE Event
	return evt, fbk
}

func newTry() *try {
	t := &try{}
	b := behavior.NewTryCmd(CMD_TOPIC, t.raiseEvent)
	t.TryCmd = b
	return t
}

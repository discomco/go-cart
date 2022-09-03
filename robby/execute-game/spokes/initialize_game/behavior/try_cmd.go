package behavior

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/robby/execute-game/-shared/specs/state_must"
	initialize_game "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/pkg/errors"
)

type ITryCmd interface {
	domain.ITryCmd
}

func TryCmd() domain.IAggPlugin {
	return newTry()
}

type try struct {
	*domain.TryCmd
}

func (t *try) raiseEvent(ctx context.Context, cmd domain.ICmd) (domain.IEvt, dtos.IFbk) {
	// Initializations
	aggID := cmd.GetAggregateID()
	fbk := dtos.NewFbk(aggID.Id(), -1, "")
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
		e := fmt.Sprint(errors.Wrapf(err, "(initialize.raiseEvent) could not extract payload"))
		fbk.SetError(e)
	}
	evt := NewEvt(agg, pl)
	// RAISE Event
	return evt, fbk
}

func newTry() *try {
	t := &try{}
	b := domain.NewTryCmd(CMD_TOPIC, t.raiseEvent)
	t.TryCmd = b
	return t
}

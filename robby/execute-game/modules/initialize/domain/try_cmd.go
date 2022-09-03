package domain

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/-shared/specs/state_must"
	initialize_dtos "github.com/discomco/go-cart/robby/execute-game/modules/initialize/dtos"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
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
	// RAISE Event
	var pl initialize_dtos.Payload
	cmd.GetJsonPayload(&pl)
	evt := NewEvt(agg, pl)
	return evt, fbk
}

func newTry() *try {
	t := &try{}
	b := domain.NewTryCmd(CMD_TOPIC, t.raiseEvent)
	t.TryCmd = b
	return t
}

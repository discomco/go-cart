package domain

import (
	"context"

	"github.com/discomco/go-cart/dtos"
)

type raiseFunc func(context.Context, ICmd) (IEvt, dtos.IFbk)

// TryCmd is the -base receiver for executor extensions
type TryCmd struct {
	aggregate IAggregate
	cmdType   CommandType
	raiseFunc raiseFunc
}

// TryCmd is the entry point for the TryCmd and is called by the composed Aggregate
func (e *TryCmd) TryCommand(ctx context.Context, command ICmd) (IEvt, dtos.IFbk) {
	return e.raiseFunc(ctx, command)
}

// GetAggregate returns the Aggregate tha this TryCmd is part of.
func (e *TryCmd) GetAggregate() IAggregate {
	return e.aggregate
}

// SetAggregate
func (e *TryCmd) SetAggregate(agg IAggregate) {
	e.aggregate = agg
}

func (e *TryCmd) GetCommandType() CommandType {
	return e.cmdType
}

// NewTryCmd returns a new Command Executor for cndType and allows you to supply an executor function.
// It automatically registers the TryCmd into the Aggregate.
func NewTryCmd(cmdType CommandType, raise raiseFunc) *TryCmd {
	result := &TryCmd{
		cmdType:   cmdType,
		raiseFunc: raise,
	}
	return result
}

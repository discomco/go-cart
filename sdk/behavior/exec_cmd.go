package behavior

import (
	"context"
	"github.com/discomco/go-cart/sdk/contract"
)

type raiseFunc func(context.Context, ICmd) (IEvt, contract.IFbk)

// TryCmd is the -base receiver for executor extensions
type TryCmd struct {
	aggregate IBehavior
	cmdType   CommandType
	raiseFunc raiseFunc
}

// TryCmd is the entry point for the TryCmd and is called by the composed Aggregate
func (e *TryCmd) TryCommand(ctx context.Context, command ICmd) (IEvt, contract.IFbk) {
	return e.raiseFunc(ctx, command)
}

// GetAggregate returns the Aggregate tha this TryCmd is part of.
func (e *TryCmd) GetAggregate() IBehavior {
	return e.aggregate
}

// SetAggregate
func (e *TryCmd) SetAggregate(agg IBehavior) {
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

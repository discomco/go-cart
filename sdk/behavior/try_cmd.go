package behavior

import (
	"context"
	"github.com/discomco/go-cart/sdk/contract"
)

type FRaise func(context.Context, ICmd) (IEvt, contract.IFbk)

// TryCmd is the -base receiver for executor extensions
type TryCmd struct {
	behavior IBehavior
	cmdType  CommandType
	fRaise   FRaise
}

// TryCommand is the entry point for the TryCmd and is called by the composed behavior
func (e *TryCmd) TryCommand(ctx context.Context, command ICmd) (IEvt, contract.IFbk) {
	return e.fRaise(ctx, command)
}

// GetBehavior returns the Aggregate tha this TryCmd is part of.
func (e *TryCmd) GetBehavior() IBehavior {
	return e.behavior
}

// SetBehavior sets the behavior of this TryCmd
func (e *TryCmd) SetBehavior(beh IBehavior) {
	e.behavior = beh
}

func (e *TryCmd) GetCommandType() CommandType {
	return e.cmdType
}

// NewTryCmd returns a new Command Executor for cndType and allows you to supply an executor function.
// It automatically registers the TryCmd into the Aggregate.
func NewTryCmd(cmdType CommandType, raise FRaise) *TryCmd {
	result := &TryCmd{
		cmdType: cmdType,
		fRaise:  raise,
	}
	return result
}

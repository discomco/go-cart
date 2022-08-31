package domain

import (
	"github.com/discomco/go-cart/sdk/dtos"
	"golang.org/x/net/context"
)

//anExec extends TryCmd and implements a raiseEvt function that contains some business logic and
type anExec struct {
	*TryCmd
}

func (e *anExec) raiseEvt(ctx context.Context, cmd ICmd) (IEvt, dtos.IFbk) {
	fbk := dtos.NewFbk(cmd.GetAggregateID().Id(), -1, "")
	evt := NewEvt(e.GetAggregate(), A_EVT_TOPIC)

	return evt, fbk
}
func newAnExecCmd() IATryCmd {
	r := &anExec{}
	exec := NewTryCmd(A_CMD_TOPIC, r.raiseEvt)
	r.TryCmd = exec
	return r
}
func AnExec() IAggPlugin {
	return newAnExecCmd()
}

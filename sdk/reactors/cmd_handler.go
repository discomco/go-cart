package reactors

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/contract"
	sdk_errors "github.com/discomco/go-cart/sdk/core/errors"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-cart/sdk/specs/allow"
	"github.com/discomco/go-cart/sdk/specs/cmd_must"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
)

const (
	CmdHandlerFmt = "%+v.CmdHandler"
)

type CmdHandlerFtor func() ICmdHandler

func CmdHandler(newAs BehSFtor, newAgg behavior.BehaviorBuilder) CmdHandlerFtor {
	return func() ICmdHandler {
		as := newAs()
		agg := newAgg()
		return newCmdHandler(agg, as)
	}
}

type cmdHandler struct {
	*Component
	topic    behavior.Topic
	behavior behavior.IBehavior
	as       IBehaviorStore
}

func (h *cmdHandler) GetAggregate(ID schema.IIdentity) behavior.IBehavior {
	if h.behavior == nil {
		panic(sdk_errors.ErrNoAggregate)
	}
	return h.behavior.SetID(ID)
}

func (h *cmdHandler) Handle(ctx context.Context, cmd behavior.ICmd) contract.IFbk {
	fbk := contract.NewFbk(cmd.GetAggregateID().Id(), -1, "")
	cmd_must.NotBeNil(cmd, fbk)
	cmd_must.HaveAggregateID(cmd, fbk)
	if !fbk.IsSuccess() {
		return fbk
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.Handle")
	defer span.Finish()
	span.LogFields(opentracing_log.String("aggregateID", cmd.GetAggregateID().Id()))

	ID := cmd.GetAggregateID()
	h.behavior.SetID(ID)

	// Do we have an Aggregate?
	err := h.as.Exists(ctx, ID.Id())

	allow.StreamNotFound(err, fbk)
	if !fbk.IsSuccess() {
		return fbk
	}

	err = h.as.Load(ctx, h.behavior)
	allow.StreamNotFound(err, fbk)
	if !fbk.IsSuccess() {
		return fbk
	}

	evt, f := h.behavior.TryCommand(ctx, cmd)
	if evt == nil {
		f.SetError(behavior.ErrExecuteDidNotReturnAnEvent(string(cmd.GetCommandType())))
		return f
	}
	span.LogFields(opentracing_log.String("domain", h.behavior.String()))
	err = h.as.Save(ctx, h.behavior)
	if err != nil {
		f.SetError(err.Error())
		return f
	}
	h.GetMediator().Broadcast(evt.GetEventTypeString(), ctx, evt)
	return f
}

func newCmdHandler(agg behavior.IBehavior, as IBehaviorStore) ICmdHandler {
	name := fmt.Sprintf(CmdHandlerFmt, agg.GetBehaviorType())
	base := NewComponent(schema.Name(name))
	h := &cmdHandler{
		Component: base,
		topic:     "",
		behavior:  agg,
		as:        as,
	}
	return h
}

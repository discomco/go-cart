package features

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core"
	sdk_errors "github.com/discomco/go-cart/sdk/core/errors"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/features/-specs/allow"
	"github.com/discomco/go-cart/sdk/features/-specs/cmd_must"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
)

const (
	CmdHandlerFmt = "%+v.CmdHandler"
)

type CmdHandlerFtor func() ICmdHandler

func CmdHandler(newAs ASFtor, newAgg domain.AggBuilder) CmdHandlerFtor {
	return func() ICmdHandler {
		as := newAs()
		agg := newAgg()
		return newCmdHandler(agg, as)
	}
}

type cmdHandler struct {
	*AppComponent
	topic domain.Topic
	agg   domain.IAggregate
	as    IAggregateStore
}

func (h *cmdHandler) GetAggregate(ID core.IIdentity) domain.IAggregate {
	if h.agg == nil {
		panic(sdk_errors.ErrNoAggregate)
	}
	return h.agg.SetID(ID)
}

func (h *cmdHandler) Handle(ctx context.Context, cmd domain.ICmd) dtos.IFbk {
	fbk := dtos.NewFbk(cmd.GetAggregateID().Id(), -1, "")
	cmd_must.NotBeNil(cmd, fbk)
	cmd_must.HaveAggregateID(cmd, fbk)
	if !fbk.IsSuccess() {
		return fbk
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, "handler.Handle")
	defer span.Finish()
	span.LogFields(opentracing_log.String("aggregateID", cmd.GetAggregateID().Id()))

	ID := cmd.GetAggregateID()
	h.agg.SetID(ID)

	// Do we have an Aggregate?
	err := h.as.Exists(ctx, ID.Id())

	allow.StreamNotFound(err, fbk)
	if !fbk.IsSuccess() {
		return fbk
	}

	err = h.as.Load(ctx, h.agg)
	allow.StreamNotFound(err, fbk)
	if !fbk.IsSuccess() {
		return fbk
	}

	evt, f := h.agg.TryCommand(ctx, cmd)
	if evt == nil {
		f.SetError(domain.ErrExecuteDidNotReturnAnEvent(string(cmd.GetCommandType())))
		return f
	}
	span.LogFields(opentracing_log.String("domain", h.agg.String()))
	err = h.as.Save(ctx, h.agg)
	if err != nil {
		f.SetError(err.Error())
		return f
	}
	h.GetMediator().Broadcast(evt.GetEventTypeString(), ctx, evt)
	return f
}

func newCmdHandler(agg domain.IAggregate, as IAggregateStore) ICmdHandler {
	name := fmt.Sprintf(CmdHandlerFmt, agg.GetAggregateType())
	base := NewAppComponent(Name(name))
	h := &cmdHandler{
		AppComponent: base,
		topic:        "",
		agg:          agg,
		as:           as,
	}
	return h
}

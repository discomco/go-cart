package comps

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type LoggingReactorFtor func() ILoggingReactor
type GenLoggingReactorFtor[TEvt behavior.IEvt] func() IGenLoggingReactor[TEvt]

type ILoggingReactor interface {
	IEvtReactor
}

type IGenLoggingReactor[TEvt behavior.IEvt] interface {
	IGenEvtReactor[TEvt]
}

type loggingReactor struct {
	*EventReactor
}

func (h *loggingReactor) handleEvent(ctx context.Context, evt behavior.IEvt) error {
	h.GetLogger().
		Debugf("[Logging Reactor for %+v] Handling domain.Evt [%+v]", evt.GetEventType(), evt)
	return nil
}

func newLoggingReactor(topic behavior.EventType) *loggingReactor {
	lh := &loggingReactor{}
	b := NewEventReactor(topic, lh.handleEvent)
	lh.EventReactor = b
	lh.Name = schema.Name(fmt.Sprintf("loggingReactor.%+v", topic))
	return lh
}

func NewLoggingReactorFtor(topic behavior.EventType) LoggingReactorFtor {
	return func() ILoggingReactor {
		return newLoggingReactor(topic)
	}
}

func NewLoggingReactor(topic behavior.EventType) ILoggingReactor {
	return newLoggingReactor(topic)
}

func GeneralMediatorLogger() ILoggingReactor {
	return newLoggingReactor(behavior.AllTopics)
}

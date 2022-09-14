package comps

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type MediatorLoggerFtor func() IMediatorLogger
type GenMediatorLoggerFtor[TEvt behavior.IEvt] func() IGenMediatorLogger[TEvt]

type IMediatorLogger interface {
	IEvtReaction
}

type IGenMediatorLogger[TEvt behavior.IEvt] interface {
	IGenEvtReaction[TEvt]
}

type mediatorLogger struct {
	*EventReaction
}

func (h *mediatorLogger) logEvent(ctx context.Context, evt behavior.IEvt) error {
	h.GetLogger().
		Debugf("[Logging for %+v] Handling domain.Evt [%+v]", evt.GetEventType(), evt)
	return nil
}

func newMediatorLogger(topic behavior.EventType) *mediatorLogger {
	lh := &mediatorLogger{}
	b := NewEventReaction(topic, lh.logEvent)
	lh.EventReaction = b
	lh.Name = schema.Name(fmt.Sprintf("mediatorLogger.%+v", topic))
	return lh
}

func NewMediatorLoggerFtor(topic behavior.EventType) MediatorLoggerFtor {
	return func() IMediatorLogger {
		return newMediatorLogger(topic)
	}
}

func NewMediatorLogger(topic behavior.EventType) IMediatorLogger {
	return newMediatorLogger(topic)
}

func GeneralMediatorLogger() IMediatorLogger {
	return newMediatorLogger(behavior.AllTopics)
}

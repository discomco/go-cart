package features

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/domain"
)

type MediatorLoggerFtor func() IMediatorLogger
type GenMediatorLoggerFtor[TEvt domain.IEvt] func() IGenMediatorLogger[TEvt]

type IMediatorLogger interface {
	IEvtHandler
}

type IGenMediatorLogger[TEvt domain.IEvt] interface {
	IGenEvtHandler[TEvt]
}

type mediatorLogger struct {
	*EventHandler
}

func (h *mediatorLogger) handleEvent(ctx context.Context, evt domain.IEvt) error {
	h.GetLogger().
		Debugf("[Logging Handler for %+v] Handling domain.Evt [%+v]", evt.GetEventType(), evt)
	return nil
}

func newMediatorLogger(topic domain.EventType) *mediatorLogger {
	lh := &mediatorLogger{}
	b := NewEventHandler(topic, lh.handleEvent)
	lh.EventHandler = b
	lh.Name = Name(fmt.Sprintf("mediatorLogger.%+v", topic))
	return lh
}

func NewMediatorLoggerFtor(topic domain.EventType) MediatorLoggerFtor {
	return func() IMediatorLogger {
		return newMediatorLogger(topic)
	}
}

func NewMediatorLogger(topic domain.EventType) IMediatorLogger {
	return newMediatorLogger(topic)
}

func GeneralMediatorLogger() IMediatorLogger {
	return newMediatorLogger(domain.AllTopics)
}

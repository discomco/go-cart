package features

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"sync"
)

type OnEvtFunc func(ctx context.Context, evt domain.IEvt) error

type EvtHandlerFtor func() IEvtHandler
type GenEvtHandlerFtor[TEvt domain.IEvt] func() IGenEvtHandler[TEvt]

type EventHandler struct {
	*AppComponent
	mediator  mediator.IMediator
	evtType   domain.EventType
	onEvt     OnEvtFunc
	whenMutex *sync.Mutex
}

const EventHandlerFmt = "%+v.EvtHandler"

func NewEventHandler(
	eventType domain.EventType,
	onEvt OnEvtFunc,
) *EventHandler {
	name := fmt.Sprintf(EventHandlerFmt, eventType)
	base := NewAppComponent(Name(name))
	result := &EventHandler{
		AppComponent: base,
		evtType:      eventType,
		onEvt:        onEvt,
		whenMutex:    &sync.Mutex{},
	}
	return result
}

func (h *EventHandler) Deactivate(ctx context.Context) error {
	if h.evtType == domain.AllTopics {
		h.UnsubscribeAll(h.When)
		return nil
	}
	err := h.unsub(string(h.evtType), h.When)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	return nil

}

func (h *EventHandler) Unsubscribe(topic string, fn OnEvtFunc) error {
	return h.unsub(topic, fn)
}

func (h *EventHandler) unsub(topic string, fn interface{}) error {
	err := h.GetMediator().Unregister(topic, fn)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	h.GetLogger().Infof("[%s] unlinked [%+v]", h.GetName(), h.GetEventType())
	return nil
}

func (h *EventHandler) GetEventType() domain.EventType {
	return h.evtType
}

func (h *EventHandler) When(ctx context.Context, evt domain.IEvt) error {
	if h.onEvt == nil {
		return nil
	}
	h.whenMutex.Lock()
	defer h.whenMutex.Unlock()
	h.GetLogger().Infof("[%+v] received [%+v]", h.GetName(), evt.GetEventType())
	return h.onEvt(ctx, evt)
}

func (h *EventHandler) Activate(ctx context.Context) error {
	if h.evtType == domain.AllTopics {
		h.SubscribeAll(ctx, h.When, true)
		return nil
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, string(h.evtType), h.When, true))
	return wg.Wait()
}

func (h *EventHandler) subWorker(ctx context.Context, topic string, fn interface{}, transactional bool) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			h.Deactivate(ctx)
			return ctx.Err()
		default:
			err := h.GetMediator().RegisterAsync(topic, fn, transactional)
			if err != nil {
				h.GetLogger().Error(err)
				return errors.Wrap(err, "failed to register with mediator")
			}
			h.GetMediator().WaitAsync()
			h.GetLogger().Infof("[%+v] links [%+v]", h.GetName(), h.evtType)
			return nil
		}
	}
}

func (h *EventHandler) SubscribeAll(ctx context.Context, when OnEvtFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		wg.Go(h.subWorker(ctx, topic, when, transactional))
	}
	return wg.Wait()
}

func (h *EventHandler) Subscribe(ctx context.Context, topic string, when OnEvtFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, topic, when, transactional))
	return wg.Wait()
}

func (h *EventHandler) SubscribeAllAsync(events chan domain.IEvt, transactional bool) map[string]error {
	res := make(map[string]error)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		err := h.subAsync(topic, events, transactional)
		if err != nil {
			h.GetLogger().Errorf("[%+v] failed to link [%s]. error: %v", h.GetName(), topic, err)
			res[topic] = err
		}
	}
	return res
}

func (h *EventHandler) subAsync(topic string, events chan domain.IEvt, transactional bool) error {
	err := h.GetMediator().RegisterAsync(topic, func(evt domain.IEvt) {
		events <- evt
	}, transactional)
	if err != nil {
		h.GetLogger().Fatal(err)
		return err
	}
	h.GetMediator().WaitAsync()
	return nil
}

func (h *EventHandler) SubscribeAsync(events chan domain.IEvt, transactional bool) error {
	return h.subAsync(string(h.evtType), events, transactional)
}

func (h *EventHandler) UnsubscribeAll(when OnEvtFunc) map[string]error {
	errs := make(map[string]error, 0)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		err := h.unsub(topic, when)
		if err != nil {
			errs[topic] = err
		}
	}
	return errs
}

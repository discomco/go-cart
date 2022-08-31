package features

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/domain"
	"golang.org/x/net/context"
)

type GenOnEvtFunc[TEvt domain.IEvt] func(ctx context.Context, evt TEvt) error

type GenEvtHandler[TEvt domain.IEvt] struct {
	*AppComponent
	evtType domain.EventType
	onEvt   OnEvtFunc
}

func (h *GenEvtHandler[TEvt]) Deactivate(ctx context.Context) error {
	err := h.GetMediator().Unregister(string(h.evtType), h.When)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	h.GetLogger().Debugf("[%+v] Deactivated.", h.GetName())
	return nil
}

func (h *GenEvtHandler[TEvt]) GetEventType() domain.EventType {
	return h.evtType
}

func (h *GenEvtHandler[TEvt]) When(ctx context.Context, evt TEvt) error {
	return h.onEvt(ctx, evt)
}

func (h *GenEvtHandler[TEvt]) Activate(ctx context.Context) error {
	select {
	case <-ctx.Done():
		{
			err := h.Deactivate(ctx)
			if err != nil {
				return err
			}
			return ctx.Err()
		}
	default:
		err := h.GetMediator().RegisterAsync(string(h.evtType), h.When, false)
		if err != nil {
			h.GetLogger().Error(err)
			return err
		}
		h.GetMediator().WaitAsync()
		h.GetLogger().Debugf("Handling domain.Evt [%+v]", h.evtType)
		return nil
	}
}

func (h *GenEvtHandler[TEvt]) SubscribeAsync(events chan TEvt, transactional bool) error {
	err := h.GetMediator().RegisterAsync(string(h.evtType), func(evt TEvt) {
		events <- evt
	}, transactional)
	if err != nil {
		h.GetLogger().Fatal(err)
		return err
	}
	h.GetMediator().WaitAsync()
	return nil
}

func newGenEvtHandler[TEvt domain.IEvt](
	eventType domain.EventType,
	onEvt OnEvtFunc,
) *GenEvtHandler[TEvt] {
	name := fmt.Sprintf(EventHandlerFmt, eventType)
	base := NewAppComponent(Name(name))
	result := &GenEvtHandler[TEvt]{
		AppComponent: base,
		evtType:      eventType,
		onEvt:        onEvt,
	}
	return result
}

func NewGenEvtHandler[TEvt domain.IEvt](eventType domain.EventType,
	onEvt OnEvtFunc) IGenEvtHandler[TEvt] {
	return newGenEvtHandler[TEvt](eventType, onEvt)
}

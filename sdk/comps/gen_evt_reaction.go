package comps

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/net/context"
)

type GenReactFunc[TEvt behavior.IEvt] func(ctx context.Context, evt TEvt) error

type GenEvtReactor[TEvt behavior.IEvt] struct {
	*Component
	evtType behavior.EventType
	react   OnEvtFunc
}

func (er *GenEvtReactor[TEvt]) Deactivate(ctx context.Context) error {
	err := er.GetMediator().Unregister(string(er.evtType), er.React)
	if err != nil {
		er.GetLogger().Error(err)
		return err
	}
	er.GetLogger().Debugf("[%+v] Deactivated.", er.GetName())
	return nil
}

func (er *GenEvtReactor[TEvt]) GetEventType() behavior.EventType {
	return er.evtType
}

func (er *GenEvtReactor[TEvt]) React(ctx context.Context, evt TEvt) error {
	return er.react(ctx, evt)
}

func (er *GenEvtReactor[TEvt]) Activate(ctx context.Context) error {
	select {
	case <-ctx.Done():
		{
			err := er.Deactivate(ctx)
			if err != nil {
				return err
			}
			return ctx.Err()
		}
	default:
		err := er.GetMediator().RegisterAsync(string(er.evtType), er.React, false)
		if err != nil {
			er.GetLogger().Error(err)
			return err
		}
		er.GetMediator().WaitAsync()
		er.GetLogger().Debugf("Handling domain.Evt [%+v]", er.evtType)
		return nil
	}
}

func (er *GenEvtReactor[TEvt]) SubscribeAsync(events chan TEvt, transactional bool) error {
	err := er.GetMediator().RegisterAsync(string(er.evtType), func(evt TEvt) {
		events <- evt
	}, transactional)
	if err != nil {
		er.GetLogger().Fatal(err)
		return err
	}
	er.GetMediator().WaitAsync()
	return nil
}

func newGenEvtHandler[TEvt behavior.IEvt](
	eventType behavior.EventType,
	onEvt OnEvtFunc,
) *GenEvtReactor[TEvt] {
	name := fmt.Sprintf(EventReactionFmt, eventType)
	base := NewComponent(schema.Name(name))
	result := &GenEvtReactor[TEvt]{
		Component: base,
		evtType:   eventType,
		react:     onEvt,
	}
	return result
}

func NewGenEvtHandler[TEvt behavior.IEvt](eventType behavior.EventType,
	onEvt OnEvtFunc) IGenEvtReaction[TEvt] {
	return newGenEvtHandler[TEvt](eventType, onEvt)
}

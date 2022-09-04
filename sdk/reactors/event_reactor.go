package reactors

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"sync"
)

type OnEvtFunc func(ctx context.Context, evt behavior.IEvt) error

type EvtReactorFtor func() IEvtReactor
type GenEvtReactorFtor[TEvt behavior.IEvt] func() IGenEvtReactor[TEvt]

type EventReactor struct {
	*Component
	mediator  mediator.IMediator
	evtType   behavior.EventType
	react     OnEvtFunc
	whenMutex *sync.Mutex
}

const EventReactorFmt = "%+v.EvtReactor"

func NewEventReactor(
	eventType behavior.EventType,
	react OnEvtFunc,
) *EventReactor {
	name := fmt.Sprintf(EventReactorFmt, eventType)
	base := NewComponent(schema.Name(name))
	result := &EventReactor{
		Component: base,
		evtType:   eventType,
		react:     react,
		whenMutex: &sync.Mutex{},
	}
	return result
}

func (er *EventReactor) Deactivate(ctx context.Context) error {
	if er.evtType == behavior.AllTopics {
		er.UnsubscribeAll(er.React)
		return nil
	}
	err := er.unsub(string(er.evtType), er.React)
	if err != nil {
		er.GetLogger().Error(err)
		return err
	}
	return nil

}

func (er *EventReactor) Unsubscribe(topic string, fn OnEvtFunc) error {
	return er.unsub(topic, fn)
}

func (er *EventReactor) unsub(topic string, fn interface{}) error {
	err := er.GetMediator().Unregister(topic, fn)
	if err != nil {
		er.GetLogger().Error(err)
		return err
	}
	er.GetLogger().Infof("[%s] unlinked [%+v]", er.GetName(), er.GetEventType())
	return nil
}

func (er *EventReactor) GetEventType() behavior.EventType {
	return er.evtType
}

func (er *EventReactor) React(ctx context.Context, evt behavior.IEvt) error {
	if er.react == nil {
		return nil
	}
	er.whenMutex.Lock()
	defer er.whenMutex.Unlock()
	er.GetLogger().Infof("[%+v] received [%+v]", er.GetName(), evt.GetEventType())
	return er.react(ctx, evt)
}

func (er *EventReactor) Activate(ctx context.Context) error {
	if er.evtType == behavior.AllTopics {
		er.SubscribeAll(ctx, er.React, true)
		return nil
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(er.subWorker(ctx, string(er.evtType), er.React, true))
	return wg.Wait()
}

func (er *EventReactor) subWorker(ctx context.Context, topic string, fn interface{}, transactional bool) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			er.Deactivate(ctx)
			return ctx.Err()
		default:
			err := er.GetMediator().RegisterAsync(topic, fn, transactional)
			if err != nil {
				er.GetLogger().Error(err)
				return errors.Wrap(err, "failed to register with mediator")
			}
			er.GetMediator().WaitAsync()
			er.GetLogger().Infof("[%+v] links [%+v]", er.GetName(), er.evtType)
			return nil
		}
	}
}

func (er *EventReactor) SubscribeAll(ctx context.Context, when OnEvtFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	topics := er.GetMediator().KnownTopics()
	for topic := range topics {
		wg.Go(er.subWorker(ctx, topic, when, transactional))
	}
	return wg.Wait()
}

func (er *EventReactor) Subscribe(ctx context.Context, topic string, when OnEvtFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(er.subWorker(ctx, topic, when, transactional))
	return wg.Wait()
}

func (er *EventReactor) SubscribeAllAsync(events chan behavior.IEvt, transactional bool) map[string]error {
	res := make(map[string]error)
	topics := er.GetMediator().KnownTopics()
	for topic := range topics {
		err := er.subAsync(topic, events, transactional)
		if err != nil {
			er.GetLogger().Errorf("[%+v] failed to link [%s]. error: %v", er.GetName(), topic, err)
			res[topic] = err
		}
	}
	return res
}

func (er *EventReactor) subAsync(topic string, events chan behavior.IEvt, transactional bool) error {
	err := er.GetMediator().RegisterAsync(topic, func(evt behavior.IEvt) {
		events <- evt
	}, transactional)
	if err != nil {
		er.GetLogger().Fatal(err)
		return err
	}
	er.GetMediator().WaitAsync()
	return nil
}

func (er *EventReactor) SubscribeAsync(events chan behavior.IEvt, transactional bool) error {
	return er.subAsync(string(er.evtType), events, transactional)
}

func (er *EventReactor) UnsubscribeAll(when OnEvtFunc) map[string]error {
	errs := make(map[string]error, 0)
	topics := er.GetMediator().KnownTopics()
	for topic := range topics {
		err := er.unsub(topic, when)
		if err != nil {
			errs[topic] = err
		}
	}
	return errs
}

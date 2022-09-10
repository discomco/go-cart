package comps

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

type EvtReactionFtor func() IEvtReaction
type GenEvtReactionFtor[TEvt behavior.IEvt] func() IGenEvtReaction[TEvt]

// EventReaction is a base structure for implementing specialized Reactions to Events
type EventReaction struct {
	*Component
	mediator  mediator.IMediator
	evtType   behavior.EventType
	react     OnEvtFunc
	whenMutex *sync.Mutex
}

const EventReactionFmt = "%+v.EvtReaction"

func NewEventReaction(
	eventType behavior.EventType,
	react OnEvtFunc,
) *EventReaction {
	name := fmt.Sprintf(EventReactionFmt, eventType)
	base := NewComponent(schema.Name(name))
	result := &EventReaction{
		Component: base,
		evtType:   eventType,
		react:     react,
		whenMutex: &sync.Mutex{},
	}
	return result
}

func (er *EventReaction) Deactivate(ctx context.Context) error {
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

func (er *EventReaction) Unsubscribe(topic string, fn OnEvtFunc) error {
	return er.unsub(topic, fn)
}

func (er *EventReaction) unsub(topic string, fn interface{}) error {
	err := er.GetMediator().Unregister(topic, fn)
	if err != nil {
		er.GetLogger().Error(err)
		return err
	}
	er.GetLogger().Infof("[%s] unlinked [%+v]", er.GetName(), er.GetEventType())
	return nil
}

func (er *EventReaction) GetEventType() behavior.EventType {
	return er.evtType
}

// React is the function that is responsible for handling events.
func (er *EventReaction) React(ctx context.Context, evt behavior.IEvt) error {
	if er.react == nil {
		return nil
	}
	er.whenMutex.Lock()
	defer er.whenMutex.Unlock()
	er.GetLogger().Infof("[%+v] received [%+v]", er.GetName(), evt.GetEventType())
	return er.react(ctx, evt)
}

// Activate activates the Reaction and subscribes to the specific Event type on the Mediator
func (er *EventReaction) Activate(ctx context.Context) error {
	if er.evtType == behavior.AllTopics {
		er.SubscribeAll(ctx, er.React, true)
		return nil
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(er.reactWorker(ctx, string(er.evtType), er.React, true))
	return wg.Wait()
}

func (er *EventReaction) reactWorker(ctx context.Context, topic string, fn interface{}, transactional bool) func() error {
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

// SubscribeAll subscribes to all events that appear on the Mediator
func (er *EventReaction) SubscribeAll(ctx context.Context, when OnEvtFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	topics := er.GetMediator().KnownTopics()
	for topic := range topics {
		wg.Go(er.reactWorker(ctx, topic, when, transactional))
	}
	return wg.Wait()
}

// Subscribe connects the Reaction to a specific Event that appears on the Mediator
func (er *EventReaction) Subscribe(ctx context.Context, topic string, when OnEvtFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(er.reactWorker(ctx, topic, when, transactional))
	return wg.Wait()
}

func (er *EventReaction) SubscribeAllAsync(events chan behavior.IEvt, transactional bool) map[string]error {
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

func (er *EventReaction) subAsync(topic string, events chan behavior.IEvt, transactional bool) error {
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

func (er *EventReaction) SubscribeAsync(events chan behavior.IEvt, transactional bool) error {
	return er.subAsync(string(er.evtType), events, transactional)
}

// UnsubscribeAll disconnects the Reaction from all topics.
func (er *EventReaction) UnsubscribeAll(when OnEvtFunc) map[string]error {
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

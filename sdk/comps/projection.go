package comps

import (
	"context"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"sync"
)

type ProjFtor[TEvt behavior.IEvt, TState schema.IReadModel] func() IProjection
type GenProjFtor[TEvt behavior.IEvt, TState schema.IReadModel] func() IGenProjection[TEvt, TState]

type IProjection interface {
	IMediatorReaction
	IAmProjection()
}

type GenProjection[TEvt behavior.IEvt, TState schema.IReadModel] struct {
	*EventReaction
	store     behavior.IReadModelStore[TState]
	evt2Doc   behavior.Evt2ModelFunc[TEvt, TState]
	newDoc    schema.DocFtor[TState]
	getDocKey schema.GetDocKeyFunc
}

var cMutex = &sync.Mutex{}

func newGenProjection[TEvt behavior.IEvt, TState schema.IReadModel](
	name schema.Name,
	eventType behavior.EventType,
	store behavior.IReadModelStore[TState],
	evt2doc behavior.Evt2ModelFunc[TEvt, TState],
	newDoc schema.DocFtor[TState],
	getDocKey schema.GetDocKeyFunc,
) *GenProjection[TEvt, TState] {
	h := &GenProjection[TEvt, TState]{
		evt2Doc:   evt2doc,
		newDoc:    newDoc,
		store:     store,
		getDocKey: getDocKey,
	}
	h.EventReaction = NewEventReaction(eventType, h.loadEvent)
	h.Name = name
	return h
}

var lMutex = &sync.Mutex{}

func (ph *GenProjection[TEvt, TState]) IAmProjection() {}

func (ph *GenProjection[TEvt, TState]) loadEvent(ctx context.Context, evt behavior.IEvt) error {
	lMutex.Lock()
	defer lMutex.Unlock()
	var doc schema.IReadModel
	key := evt.GetAggregateId()
	if ph.getDocKey != nil {
		key = ph.getDocKey()
	}
	doc, err := ph.store.Get(ctx, key)
	if err != nil {
		if err != redis.Error(redis.Nil) {
			return errors.Wrapf(err, "loadEvent: failed to get aggregate %s from cache", evt.GetAggregateId())
		}
		doc = ph.newDoc()
	}
	err = ph.evt2Doc(evt.(TEvt), doc.(*TState))
	if err != nil {
		return errors.Wrapf(err, "loadEvent: failed to map event %+v to doc %+v", evt, doc)
	}
	_, err = ph.store.Set(ctx, key, *(doc.(*TState)))
	if err != nil {
		return errors.Wrapf(err, "loadEvent: failed to set aggregate %s to cache", key)
	}
	return err
}

func NewProjection[TEvt behavior.IEvt, TState schema.IReadModel](
	name schema.Name,
	eventType behavior.EventType,
	newStore behavior.StoreFtor[TState],
	evt2Doc behavior.Evt2ModelFunc[TEvt, TState],
	newDoc schema.DocFtor[TState],
	getDocKey schema.GetDocKeyFunc) *GenProjection[TEvt, TState] {
	return newGenProjection[TEvt, TState](
		name,
		eventType,
		newStore(),
		evt2Doc,
		newDoc,
		getDocKey)
}

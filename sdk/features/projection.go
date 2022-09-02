package features

import (
	"context"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/model"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"sync"
)

type ProjFtor[TEvt domain.IEvt, TState model.IReadModel] func() IProjection
type GenProjFtor[TEvt domain.IEvt, TState model.IReadModel] func() IGenProjection[TEvt, TState]

type IProjection interface {
	IMediatorSubscriber
	IAmProjection()
}

type GenProjection[TEvt domain.IEvt, TState model.IReadModel] struct {
	*EventHandler
	store     domain.IReadModelStore[TState]
	evt2Doc   domain.Evt2ModelFunc[TEvt, TState]
	newDoc    model.DocFtor[TState]
	getDocKey model.GetDocKeyFunc
}

var cMutex = &sync.Mutex{}

func newGenProjection[TEvt domain.IEvt, TState model.IReadModel](
	name Name,
	eventType domain.EventType,
	store domain.IReadModelStore[TState],
	evt2doc domain.Evt2ModelFunc[TEvt, TState],
	newDoc model.DocFtor[TState],
	getDocKey model.GetDocKeyFunc,
) *GenProjection[TEvt, TState] {
	h := &GenProjection[TEvt, TState]{
		evt2Doc:   evt2doc,
		newDoc:    newDoc,
		store:     store,
		getDocKey: getDocKey,
	}
	h.EventHandler = NewEventHandler(eventType, h.loadEvent)
	h.Name = name
	return h
}

var lMutex = &sync.Mutex{}

func (ph *GenProjection[TEvt, TState]) IAmProjection() {}

func (ph *GenProjection[TEvt, TState]) loadEvent(ctx context.Context, evt domain.IEvt) error {
	lMutex.Lock()
	defer lMutex.Unlock()
	var doc model.IReadModel
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

func NewProjection[TEvt domain.IEvt, TState model.IReadModel](
	name Name,
	eventType domain.EventType,
	newStore domain.StoreFtor[TState],
	evt2Doc domain.Evt2ModelFunc[TEvt, TState],
	newDoc model.DocFtor[TState],
	getDocKey model.GetDocKeyFunc) *GenProjection[TEvt, TState] {
	return newGenProjection[TEvt, TState](
		name,
		eventType,
		newStore(),
		evt2Doc,
		newDoc,
		getDocKey)
}

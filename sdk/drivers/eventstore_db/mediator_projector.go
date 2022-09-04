package eventstore_db

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/constants"
	"github.com/discomco/go-cart/sdk/drivers/convert"
	"github.com/discomco/go-cart/sdk/drivers/jaeger"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ProjectionWorkerFunc func(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error
type SubscriptionGroupName string
type ProjectorName string

//EvtProjFtor is a functor for Projectors that marshal events onto the Mediator
//where they can be handled by all kinds of GenProjection Handlers.
func EvtProjFtor(newClient EventStoreDBFtor) comps.ProjectorFtor {
	return func() comps.IProjector {
		clt := newClient()
		return newProjector(clt)
	}
}

var (
	singletonProjector comps.IProjector
	cPMutex            = &sync.Mutex{}
)

//EventProjector is a Projector that marshals events onto the Mediator
func EventProjector(newClient EventStoreDBFtor) comps.IProjector {
	if singletonProjector == nil {
		cPMutex.Lock()
		defer cPMutex.Unlock()
		clt := newClient()
		singletonProjector = newProjector(clt)
	}
	return singletonProjector
}

type eventProjector struct {
	*comps.Component
	handlers    map[behavior.EventType]comps.IProjection
	esdb        *esdb.Client
	handleMutex *sync.Mutex
}

func (o *eventProjector) Activate(ctx context.Context) error {
	//	o.GetLogger().Infof("Activating eventProjector [%+v] with %+v handlers", o.Name, len(o.handlers))
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			o.GetLogger().Errorf("Mitigating (Activate.processStream) panic: {%v}", r)
			o.GetLogger().Infof("Reactivating EventProjector [%+v] with %+v handlers", o.Name, len(o.handlers))
			o.Activate(ctx)
		}
	}()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(o.subscriberFunc(ctx, cancel))
	return g.Wait()
}

func (o *eventProjector) subscriberFunc(ctx context.Context, cancel context.CancelFunc) func() error {
	cfg := o.GetConfig().GetProjectionConfig()
	return func() error {
		err := o.Project(ctx, []string{cfg.GetEventPrefix()}, cfg.GetPoolSize())
		if err != nil {
			o.GetLogger().Errorf("EventStoreDB.Projector [%+v] error: %v", o.Name, err)
			cancel()
		}
		return err
	}
}

func (o *eventProjector) Deactivate(ctx context.Context) error {
	cfg := o.GetConfig().GetProjectionConfig()
	o.GetLogger().Infof("Deactivating eventProjector [%+v] for Subscription Group [%+v]", o.Name, cfg.GetGroup())
	o.GetESDB().Close()
	return nil
}

const (
	ProjectorFmt = "ESDBProjector"
)

func newProjector(
	esdb *esdb.Client,
) *eventProjector {
	name := ProjectorFmt
	base := comps.NewComponent(schema.Name(name))
	base.Name = "eventstoreDB.Projector"
	p := &eventProjector{
		esdb:        esdb,
		handlers:    make(map[behavior.EventType]comps.IProjection),
		handleMutex: &sync.Mutex{},
	}
	p.Component = base
	return p
}

func (o *eventProjector) GetESDB() *esdb.Client {
	return o.esdb
}

func (o *eventProjector) runWorker(ctx context.Context, worker ProjectionWorkerFunc, stream *esdb.PersistentSubscription, i int) func() error {
	return func() error {
		return worker(ctx, stream, i)
	}
}

func (o *eventProjector) Inject(projections ...comps.IProjection) {
	for _, h := range projections {
		_, ok := o.handlers[h.GetEventType()]
		if !ok {
			o.GetLogger().Debugf("(GenProjection.Inject) [%+v]", h.GetName())
			o.handlers[h.GetEventType()] = h
		}
	}
}

func (o *eventProjector) Project(ctx context.Context, prefixes []string, poolSize int) error {
	cfg := o.GetConfig().GetProjectionConfig()
	//	o.GetLogger().Infof("(starting subscription [%+v], group [%+v]) prefixes: {%+v}", cfg.GetName(), cfg.GetGroup(), prefixes)
	err := o.GetESDB().CreatePersistentSubscriptionToAll(ctx, cfg.GetGroup(), esdb.PersistentAllSubscriptionOptions{
		Filter: &esdb.SubscriptionFilter{Type: esdb.StreamFilterType, Prefixes: prefixes},
	})
	if err != nil {
		dbErr, ok := esdb.FromError(err)
		if !ok {
			switch dbErr.Code() {
			case esdb.ErrorResourceAlreadyExists:
				//				o.GetLogger().Warnf("(CreatePersistentSubscriptionAll) GenProjection [%+v] already exists.", cfg.GetGroup())
			default:
				o.GetLogger().Errorf("(CreatePersistentSubscriptionAll) err: {%v}", dbErr.Code())
			}
		}
	}

	stream, err := o.GetESDB().SubscribeToPersistentSubscription(
		ctx,
		constants.EsAll,
		cfg.GetGroup(),
		esdb.SubscribeToPersistentSubscriptionOptions{},
	)

	if err != nil {
		return err
	}
	defer stream.Close()

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= poolSize; i++ {
		g.Go(o.runWorker(ctx, o.processStream, stream, i))
	}
	return g.Wait()
}

func (o *eventProjector) processStream(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error {

	// TODO: check for context cancellation

	c := o.GetConfig().GetProjectionConfig()
	for {

		event := stream.Recv()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if event.SubscriptionDropped != nil {
				o.GetLogger().Errorf("(SubscriptionDropped) err: {%v}", event.SubscriptionDropped.Error)
				return errors.Wrap(event.SubscriptionDropped.Error, "Subscription Dropped")
			}

			if event.EventAppeared != nil {
				o.GetLogger().(Logger).ProjectionEvent(constants.MediatorProjection, c.GetGroup(), event.EventAppeared.Event, workerID)
				evt := event.EventAppeared.Event.Event

				err := o.React(ctx, convert.Recorded2Evt(evt))

				if err != nil {
					o.GetLogger().Errorf("(GenProjection.when) \nevent: {%v}, \nerr): {%v}", evt, err)
					if err := stream.Nack(err.Error(), esdb.Nack_Retry, event.EventAppeared.Event); err != nil {
						o.GetLogger().Errorf("(stream.Nack) err: {%v}", err)
						return errors.Wrap(err, "stream.Nack")
					}
				}
				err = stream.Ack(event.EventAppeared.Event)
				if err != nil {
					o.GetLogger().Errorf("(stream.Ack) err: {%v}", err)
					return errors.Wrap(err, "stream.Ack")
				}
				o.GetLogger().Infof("(ACK) event commit: {%v}", *event.EventAppeared.Event.Commit)
			}
		}
	}
}

func (o *eventProjector) React(ctx context.Context, evt behavior.IEvt) error {
	ctx, span := jaeger.StartProjectionTracerSpan(ctx, "GenProjection.React", evt)
	defer span.Finish()
	span.LogFields(
		log.String("Id [%+v]", evt.GetAggregateId()),
		log.String("EventType", evt.GetEventTypeString()))
	o.GetLogger().Debugf("(GenProjection.React) event_type: [%v], event  {%+v}", evt.GetEventType(), evt)
	topic := evt.GetEventTypeString()
	o.GetMediator().Broadcast(topic, ctx, evt)
	return nil
}

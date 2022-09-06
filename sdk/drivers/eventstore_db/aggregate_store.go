package eventstore_db

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	sdk_errors "github.com/discomco/go-cart/sdk/core/errors"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/drivers/convert"
	"github.com/discomco/go-cart/sdk/drivers/jaeger"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"io"
	"math"
	"sync"
)

const (
	count = math.MaxInt64
)

type aggregateStore struct {
	log logger.IAppLogger
	db  *esdb.Client
}

var (
	singleton interface{}
	cMutex    = &sync.Mutex{}
)

// AStore is an Injection that injects a functor for IBehaviorStore
func AStore(log logger.IAppLogger, newDb EventStoreDBFtor) comps.BehaviorStoreFtor {
	return func() comps.IBehaviorStore {
		db := newDb()
		return aStore(log, db)
	}
}

func aStore(log logger.IAppLogger, db *esdb.Client) comps.IBehaviorStore {
	return &aggregateStore{log: log, db: db}
}

func (a *aggregateStore) Load(ctx context.Context, aggregate behavior.IBehavior) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aggregateStore.Load")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", aggregate.GetID().Id()))

	stream, err := a.db.ReadStream(ctx, aggregate.GetID().Id(), esdb.ReadStreamOptions{}, count)
	if err != nil {
		jaeger.TraceErr(span, err)
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	for {
		event, err := stream.Recv()
		if err != nil {
			esdbErr, ok := esdb.FromError(err)
			if !ok && esdbErr.Code() == esdb.ErrorResourceNotFound {
				err := sdk_errors.ErrStreamNotFound
				jaeger.TraceErr(span, err)
				return errors.Wrap(err, "stream.Recv")
			}
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				jaeger.TraceErr(span, err)
				return errors.Wrap(err, "stream.Recv")
			}
		}
		esEvent := convert.Recorded2Evt(event.Event)
		if err := aggregate.RaiseEvent(esEvent); err != nil {
			jaeger.TraceErr(span, err)
			return errors.Wrap(err, "RaiseEvent")
		}
		a.log.Debugf("(Load) esEvent: {%s}", esEvent.String())
	}

	a.log.Debugf("(Load) domain: {%s}", aggregate.String())
	return nil
}

func (a *aggregateStore) Save(ctx context.Context, aggregate behavior.IBehavior) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aggregateStore.Save")
	defer span.Finish()
	span.LogFields(log.String("domain", aggregate.String()))

	if len(aggregate.GetUncommittedEvents()) == 0 {
		a.log.Debugf("(Save) [no uncommittedEvents] len: {%d}", len(aggregate.GetUncommittedEvents()))
		return nil
	}

	eventsData := make([]esdb.EventData, 0, len(aggregate.GetUncommittedEvents()))
	for _, event := range aggregate.GetUncommittedEvents() {
		eventsData = append(eventsData, convert.Evt2EventData(event))
	}

	// check for domain.GetVersion() == 0 or len(domain.GetAppliedEvents()) == 0 means new domain
	var expectedRevision esdb.ExpectedRevision
	if aggregate.GetVersion() == 0 {
		expectedRevision = esdb.NoStream{}
		a.log.Debugf("(Save) expectedRevision: {%TA}", expectedRevision)

		appendStream, err := a.db.AppendToStream(
			ctx,
			aggregate.GetID().Id(),
			esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
			eventsData...,
		)
		if err != nil {
			jaeger.TraceErr(span, err)
			return errors.Wrap(err, "db.AppendToStream")
		}

		a.log.Debugf("(Save) stream: {%+v}", appendStream)
		return nil
	}

	readOps := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.End{}}
	stream, err := a.db.ReadStream(context.Background(), aggregate.GetID().Id(), readOps, 1)
	if err != nil {
		jaeger.TraceErr(span, err)
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	lastEvent, err := stream.Recv()
	if err != nil {
		jaeger.TraceErr(span, err)
		return errors.Wrap(err, "stream.Recv")
	}

	expectedRevision = esdb.Revision(lastEvent.OriginalEvent().EventNumber)
	a.log.Debugf("(Save) expectedRevision: {%TA}", expectedRevision)

	appendStream, err := a.db.AppendToStream(
		ctx,
		aggregate.GetID().Id(),
		esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
		eventsData...,
	)
	if err != nil {
		jaeger.TraceErr(span, err)
		return errors.Wrap(err, "db.AppendToStream")
	}

	a.log.Debugf("(Save) stream: {%+v}", appendStream)
	aggregate.ClearUncommittedEvents()
	return nil
}

func (a *aggregateStore) Close() error {
	return a.db.Close()
}

// Exists checks whether the Event Stream identified by streamID exists in the EventStore.
func (a *aggregateStore) Exists(ctx context.Context, streamID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "aggregateStore.Exists")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", streamID))

	readStreamOptions := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.Revision(1)}

	stream, err := a.db.ReadStream(ctx, streamID, readStreamOptions, 1)
	if err != nil {
		return errors.Wrap(err, "db.ReadStream")
	}
	defer stream.Close()

	for {
		_, err := stream.Recv()
		if err != nil {
			esdbErr, ok := esdb.FromError(err)
			if !ok && esdbErr.Code() == esdb.ErrorResourceNotFound {
				err = sdk_errors.ErrStreamNotFound
			}
			if errors.Is(err, io.EOF) {
				break
			}
			jaeger.TraceErr(span, err)
			return errors.Wrap(err, "stream.Recv")
		}
	}
	return nil
}

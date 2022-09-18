package eventstore_db

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/comps"
	sdk_errors "github.com/discomco/go-cart/sdk/core/errors"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/drivers/convert"
	"io"

	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/drivers/jaeger"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

type eventStore struct {
	log logger.IAppLogger
	db  *esdb.Client
}

func newEventStore(log logger.IAppLogger, db *esdb.Client) comps.IEventStore {
	return &eventStore{log: log, db: db}
}

func EventStore(log logger.IAppLogger, newClient EventStoreDBFtor) comps.EventStoreFtor {
	return func() comps.IEventStore {
		db := newClient()
		return newEventStore(log, db)
	}
}

func (es *eventStore) SaveSnapshot(ctx context.Context, aggregate behavior.IBehavior) error {
	panic("not implemented")
	//streamID := aggregate.GetID()
	//span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.SaveSnapShot")
	//defer span.Finish()
	//span.LogFields(log.String("aggregateID", streamID.Id()))
	//	es.db.AppendToStream()

}

func (es *eventStore) GetSnapshot(ctx context.Context, id string) (*behavior.Snapshot, error) {
	//TODO implement me
	panic("implement me")
}

func (es *eventStore) SaveEvents(ctx context.Context, streamID string, events []behavior.IEvt) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.SaveEvents")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", streamID))

	eventsData := make([]esdb.EventData, 0, len(events))
	for _, event := range events {
		eventsData = append(eventsData, convert.Evt2EventData(event))
	}

	stream, err := es.db.AppendToStream(ctx, streamID, esdb.AppendToStreamOptions{}, eventsData...)
	if err != nil {
		jaeger.TraceErr(span, err)
		return err
	}

	es.log.Debugf("SaveEvents stream: %+v", stream)
	return nil
}

func (es *eventStore) LoadEvents(ctx context.Context, streamID string) ([]behavior.IEvt, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.Load")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", streamID))

	stream, err := es.db.ReadStream(ctx, streamID, esdb.ReadStreamOptions{
		Direction: esdb.Forwards,
		From:      esdb.Revision(1),
	}, 100)
	if err != nil {
		jaeger.TraceErr(span, err)
		return nil, err
	}
	defer stream.Close()

	events := make([]behavior.IEvt, 0, 100)
	for {
		event, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			jaeger.TraceErr(span, err)
			return nil, err
		}
		events = append(events, convert.Recorded2Evt(event.Event))
	}

	return events, nil
}

func (es *eventStore) Load(ctx context.Context, aggregate behavior.IBehavior) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "evtStoreDB.Load")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", aggregate.GetID().Id()))

	stream, err := es.db.ReadStream(ctx, aggregate.GetID().Id(), esdb.ReadStreamOptions{}, count)
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
		es.log.Debugf("(Load) esEvent: {%s}", esEvent.String())
	}

	es.log.Debugf("(Load) domain: {%s}", aggregate.String())
	return nil
}

func (es *eventStore) Save(ctx context.Context, aggregate behavior.IBehavior) error {

	span, ctx := opentracing.StartSpanFromContext(ctx, "evtStoreDB.Save")
	defer span.Finish()
	span.LogFields(log.String("behavior", aggregate.String()))

	if len(aggregate.GetUncommittedEvents()) == 0 {
		es.log.Debugf("(Save) [no uncommittedEvents] len: {%d}", len(aggregate.GetUncommittedEvents()))
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
		es.log.Debugf("(Save) expectedRevision: {%TA}", expectedRevision)

		appendStream, err := es.db.AppendToStream(
			ctx,
			aggregate.GetID().Id(),
			esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
			eventsData...,
		)
		if err != nil {
			jaeger.TraceErr(span, err)
			return errors.Wrap(err, "db.AppendToStream")
		}

		es.log.Debugf("(Save) stream: {%+v}", appendStream)
		return nil
	}

	readOps := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.End{}}
	stream, err := es.db.ReadStream(context.Background(), aggregate.GetID().Id(), readOps, 1)
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
	es.log.Debugf("(Save) expectedRevision: {%TA}", expectedRevision)

	appendStream, err := es.db.AppendToStream(
		ctx,
		aggregate.GetID().Id(),
		esdb.AppendToStreamOptions{ExpectedRevision: expectedRevision},
		eventsData...,
	)
	if err != nil {
		jaeger.TraceErr(span, err)
		return errors.Wrap(err, "db.AppendToStream")
	}

	es.log.Debugf("(Save) stream: {%+v}", appendStream)
	aggregate.ClearUncommittedEvents()
	return nil
}

func (es *eventStore) Close() error {
	return es.db.Close()
}

// Exists checks whether the Event Stream identified by streamID exists in the EventStore.
func (es *eventStore) Exists(ctx context.Context, streamID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "evtStoreDB.Exists")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", streamID))

	readStreamOptions := esdb.ReadStreamOptions{Direction: esdb.Backwards, From: esdb.Revision(1)}

	stream, err := es.db.ReadStream(ctx, streamID, readStreamOptions, 1)
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

package eventstore_db

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/drivers/convert"
	"github.com/discomco/go-cart/sdk/reactors"
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

func newEventStore(log logger.IAppLogger, db *esdb.Client) reactors.IEventStore {
	return &eventStore{log: log, db: db}
}

func EStore(log logger.IAppLogger, newClient EventStoreDBFtor) reactors.ESFtor {
	return func() reactors.IEventStore {
		db := newClient()
		return newEventStore(log, db)
	}
}

func (e *eventStore) SaveEvents(ctx context.Context, streamID string, events []behavior.IEvt) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.SaveEvents")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", streamID))

	eventsData := make([]esdb.EventData, 0, len(events))
	for _, event := range events {
		eventsData = append(eventsData, convert.Evt2EventData(event))
	}

	stream, err := e.db.AppendToStream(ctx, streamID, esdb.AppendToStreamOptions{}, eventsData...)
	if err != nil {
		jaeger.TraceErr(span, err)
		return err
	}

	e.log.Debugf("SaveEvents stream: %+v", stream)
	return nil
}

func (e *eventStore) LoadEvents(ctx context.Context, streamID string) ([]behavior.IEvt, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventStore.Load")
	defer span.Finish()
	span.LogFields(log.String("aggregateID", streamID))

	stream, err := e.db.ReadStream(ctx, streamID, esdb.ReadStreamOptions{
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

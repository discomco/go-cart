package eventstore_db

import (
	"context"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/drivers/convert"
	uuid2 "github.com/gofrs/uuid"
	"math/rand"
	"time"
)

func projectorWorker(ctx context.Context) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return testProjector.Activate(ctx)
		}
	}
}

func pusherWorker(ctx context.Context) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():

				return ctx.Err()
			default:
				eventId, _ := uuid2.NewV4()
				evt := convert.EventData2Evt(esdb.EventData{
					EventID:     eventId,
					EventType:   "event-type",
					ContentType: esdb.JsonContentType,
					Data:        nil,
					Metadata:    nil,
				})
				i := rand.Intn(2)
				err := testES.SaveEvents(ctx, testStreamIDs[i], []domain.IEvt{evt})
				if err != nil {
					testLogger.Fatal(err)
				}
				time.Sleep(5 * time.Second)
			}
		}
	}
}

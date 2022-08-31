package eventstore_db

import (
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/core/logger"
)

type Logger interface {
	logger.IAppLogger
	ProjectionEvent(projectionName string, groupName string, event *esdb.ResolvedEvent, workerID int)
}

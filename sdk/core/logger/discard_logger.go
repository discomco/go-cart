package logger

import (
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"time"
)

// Discard is a logger that is notdoing anything.
var Discard IAppLogger = new(discardLogger)

type discardLogger struct{}

func (*discardLogger) InitLogger() {}

func (*discardLogger) Sync() error { return nil }

func (*discardLogger) Debug(...interface{}) {}

func (*discardLogger) Debugf(string, ...interface{}) {}

func (*discardLogger) Info(...interface{}) {}

func (*discardLogger) Infof(string, ...interface{}) {}

func (*discardLogger) Warn(...interface{}) {}

func (*discardLogger) Warnf(string, ...interface{}) {}

func (*discardLogger) WarnMsg(string, error) {}

func (*discardLogger) Error(...interface{}) {}

func (*discardLogger) Errorf(string, ...interface{}) {}

func (*discardLogger) Err(string, error) {}

func (*discardLogger) DPanic(...interface{}) {}

func (*discardLogger) DPanicf(string, ...interface{}) {}

func (*discardLogger) Fatal(...interface{}) {}

func (*discardLogger) Fatalf(string, ...interface{}) {}

func (*discardLogger) Printf(string, ...interface{}) {}

func (*discardLogger) WithName(string) {}

func (*discardLogger) HttpMiddlewareAccessLogger(string, string, int, int64, time.Duration) {}

func (*discardLogger) GrpcMiddlewareAccessLogger(string, time.Duration, map[string][]string, error) {}

func (*discardLogger) GrpcClientInterceptorLogger(string, interface{}, interface{}, time.Duration, map[string][]string, error) {
}

func (*discardLogger) KafkaProcessMessage(string, int, string, int, int64, time.Time) {}

func (*discardLogger) KafkaLogCommittedMessage(string, int, int64) {}

func (*discardLogger) ProjectionEvent(string, string, *esdb.ResolvedEvent, int) {}

package logger

import (
	"context"
)

type ctxLogger struct{}

// WithLogger embeds logger into the drivers.
func WithLogger(ctx context.Context, log IAppLogger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, log)
}

// GetLogger retrieves logger from the drivers. It always returns logger instance.
func GetLogger(ctx context.Context) IAppLogger {
	if ret, ok := ctx.Value(ctxLogger{}).(IAppLogger); ok {
		return ret
	}
	return Discard
}

package http

import "context"

type IEndpoint interface {
	Get(ctx context.Context, path string, params map[string]string) ([]byte, error)
	Post(ctx context.Context, path string, params map[string]string, body []byte) ([]byte, error)
	Put(ctx context.Context, path string, params map[string]string, body []byte) ([]byte, error)
	Delete(ctx context.Context, path string, params map[string]string) ([]byte, error)
}

type Endpoint struct {
}

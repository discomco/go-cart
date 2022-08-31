package domain

import (
	"github.com/discomco/go-cart/sdk/model"
	"golang.org/x/net/context"
)

// StoreFtor of Type IReadModel is a functor type for functions that return an IReadModelStore that returns a
type StoreFtor[T model.IReadModel] func() IReadModelStore[T]

// IStore is an untyped Injector for the cache
type IStore interface{}

// IReadModelStore is the Injector for a Store that is discriminated by the Read-Model Type Injector
type IReadModelStore[T model.IReadModel] interface {
	IStore
	Exists(ctx context.Context, key string) (bool, error)
	Get(ctx context.Context, key string) (*T, error)
	Set(ctx context.Context, key string, data T) (string, error)
	Delete(ctx context.Context, key string) (*T, error)
}

package behavior

import (
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/net/context"
)

// StoreFtor of Type ISchema is a functor type for functions that return an IModelStore that returns a
type StoreFtor[T schema.ISchema] func() IModelStore[T]

// IStore is an untyped Injector for the cache
type IStore interface{}

// IModelStore is the Injector for a Store that is discriminated by the Read-Model Type Injector
type IModelStore[T schema.ISchema] interface {
	IStore
	Exists(ctx context.Context, key string) (bool, error)
	Get(ctx context.Context, key string) (*T, error)
	Set(ctx context.Context, key string, data T) (string, error)
	Delete(ctx context.Context, key string) (*T, error)
}

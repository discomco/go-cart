package redis

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/drivers/redis"
)

type IDocStore interface {
	behavior.IModelStore[schema.QuadraticDoc]
}

func newDocStore(cfg config.IAppConfig) IDocStore {
	newStore := redis.NewRedisStore[schema.QuadraticDoc](cfg)
	return newStore()
}

func DocStore(config config.IAppConfig) behavior.StoreFtor[schema.QuadraticDoc] {
	return func() behavior.IModelStore[schema.QuadraticDoc] {
		return newDocStore(config)
	}
}

func GetDoc(ctx context.Context, store IDocStore, key string) (*schema.QuadraticDoc, error) {
	return store.Get(ctx, key)
}

func SetDoc(ctx context.Context, store IDocStore, key string, doc *schema.QuadraticDoc) (string, error) {
	return store.Set(ctx, key, *doc)
}

package redis

import (
	"context"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/drivers/redis"
)

type IDocStore interface {
	behavior.IModelStore[schema.GameDoc]
}

func newDocStore(cfg config.IAppConfig) IDocStore {
	newStore := redis.NewRedisStore[schema.GameDoc](cfg)
	return newStore()
}

func DocStore(config config.IAppConfig) behavior.StoreFtor[schema.GameDoc] {
	return func() behavior.IModelStore[schema.GameDoc] {
		return newDocStore(config)
	}
}

func GetDoc(ctx context.Context, store IDocStore, key string) (*schema.GameDoc, error) {
	return store.Get(ctx, key)
}

func SetDoc(ctx context.Context, store IDocStore, key string, doc *schema.GameDoc) (string, error) {
	return store.Set(ctx, key, *doc)
}

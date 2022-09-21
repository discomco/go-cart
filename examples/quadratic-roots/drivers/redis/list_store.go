package redis

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/drivers/redis"
)

type IListStore interface {
	behavior.IModelStore[schema.QuadraticList]
}

func newListStore(cfg config.IAppConfig) IListStore {
	newStore := redis.NewRedisStore[schema.QuadraticList](cfg)
	return newStore()
}

func ListStore(config config.IAppConfig) behavior.StoreFtor[schema.QuadraticList] {
	return func() behavior.IModelStore[schema.QuadraticList] {
		return newListStore(config)
	}
}

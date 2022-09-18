package redis

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/drivers/redis"
)

type IListStore interface {
	behavior.IReadModelStore[schema.GameList]
}

func newListStore(cfg config.IAppConfig) IListStore {
	newStore := redis.NewRedisStore[schema.GameList](cfg)
	return newStore()
}

func ListStore(config config.IAppConfig) behavior.StoreFtor[schema.GameList] {
	return func() behavior.IReadModelStore[schema.GameList] {
		return newListStore(config)
	}
}

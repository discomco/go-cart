package redis

import (
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/drivers/redis"
)

type IListStore interface {
	domain.IReadModelStore[schema.GameList]
}

func newListStore(cfg config.IAppConfig) IListStore {
	newStore := redis.NewRedisStore[schema.GameList](cfg)
	return newStore()
}

func ListStore(config config.IAppConfig) domain.StoreFtor[schema.GameList] {
	return func() domain.IReadModelStore[schema.GameList] {
		return newListStore(config)
	}
}

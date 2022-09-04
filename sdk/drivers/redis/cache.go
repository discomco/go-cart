package redis

import (
	"encoding/json"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/go-redis/redis/v9"
	"golang.org/x/net/context"
	"log"
	"sync"
)

var cMutex = &sync.Mutex{}
var singleton interface{}

type cache[T schema.IReadModel] struct {
	client *redis.Client
}

func (c *cache[T]) Exists(ctx context.Context, key string) (bool, error) {
	cmd := c.client.Exists(ctx, key)
	return cmd.Val() != 0, cmd.Err()
}

func (c *cache[T]) Get(ctx context.Context, key string) (*T, error) {
	cmd := c.client.Get(ctx, key)
	data, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	var result T
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *cache[T]) Set(ctx context.Context, key string, value T) (string, error) {
	v, err := json.Marshal(value)
	if err != nil {
		return "NOK", err
	}
	cmd := c.client.Set(ctx, key, v, 0)
	return cmd.Result()
}

func (c *cache[T]) Delete(ctx context.Context, key string) (*T, error) {
	ref, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	cmd := c.client.Del(ctx, key)
	_, err = cmd.Result()
	if err != nil {
		return nil, err
	}
	return ref, nil
}

func newRedis[T schema.IReadModel](cfg config.IAppConfig) (behavior.IReadModelStore[T], error) {
	c := &cache[T]{}
	opts, err := redis.ParseURL(cfg.GetRedisConfig().GetUrl())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	c.client = redis.NewClient(opts)
	return c, nil
}

func oneRedis[T schema.IReadModel](cfg config.IAppConfig) (behavior.IReadModelStore[T], error) {
	cMutex.Lock()
	defer cMutex.Unlock()
	if singleton == nil {
		s, err := newRedis[T](cfg)
		if err != nil {
			return nil, err
		}
		singleton = s
	}
	return singleton.(*cache[T]), nil
}

func NewRedisStore[T schema.IReadModel](config config.IAppConfig) behavior.StoreFtor[T] {
	return func() behavior.IReadModelStore[T] {
		c, err := newRedis[T](config)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		return c
	}
}

func SingleRedisStore[T schema.IReadModel](config config.IAppConfig) behavior.StoreFtor[T] {
	return func() behavior.IReadModelStore[T] {
		c, err := oneRedis[T](config)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		return c
	}
}

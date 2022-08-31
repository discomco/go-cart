package mongo_db

import (
	"context"
	"github.com/discomco/go-cart/sdk/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
	"time"
)

type MongoDbClientFtor func() *mongo.Client

func newMongoDb(cfg config.IMongoDbConfig) (*mongo.Client, error) {
	credentials := options.Credential{
		Username:      cfg.GetUser(),
		Password:      cfg.GetPassword(),
		AuthMechanism: cfg.GetAuthMechanism(),
	}
	clientOpts := options.
		Client().
		ApplyURI(cfg.GetUri()).
		SetAuth(credentials)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		cancel()
		return nil, err
	}
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		cancel()
		return nil, err
	}
	return c, err
}

var (
	singleton *mongo.Client
	cMutex    = &sync.Mutex{}
)

func oneMongoDb(cfg config.IMongoDbConfig) (*mongo.Client, error) {
	if singleton == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		var err error
		singleton, err = newMongoDb(cfg)
		if err != nil {
			return nil, err
		}
		return singleton, nil
	}
	return singleton, nil
}

//TransientMongoClient returns a new MongoClient every time it is called.
func TransientMongoClient(appCfg config.IAppConfig) MongoDbClientFtor {
	cfg := appCfg.GetMongoDbConfig()
	return func() *mongo.Client {
		clt, err := newMongoDb(cfg)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		return clt
	}
}

//SingletonMongoClient is the functor that returns a singleton MongoDB client
func SingletonMongoClient(appCfg config.IAppConfig) MongoDbClientFtor {
	cfg := appCfg.GetMongoDbConfig()
	return func() *mongo.Client {
		clt, err := oneMongoDb(cfg)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		return clt
	}
}

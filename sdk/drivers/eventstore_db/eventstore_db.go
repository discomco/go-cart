package eventstore_db

import (
	"fmt"
	"log"
	"sync"

	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/core/errors"
)

// EventStoreDBFtor is a functor that returns an EventStoreDB client
type EventStoreDBFtor func() *esdb.Client

const (
	NoESDB = "No ESDB"
)

var (
	ErrNoESDB = fmt.Errorf(NoESDB)
)

var (
	singleClient *esdb.Client
	ccMutex      = &sync.Mutex{}
)

func transientESClient(cfg config.IESDBConfig) (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(cfg.GetConnectionString())
	if err != nil {
		return nil, err
	}
	return esdb.NewClient(settings)
}

func singletonESClient(cfg config.IESDBConfig) (*esdb.Client, error) {
	if singleClient == nil {
		ccMutex.Lock()
		defer ccMutex.Unlock()
		var err error
		singleClient, err = transientESClient(cfg)
		if err != nil {
			return nil, err
		}
		return singleClient, nil
	}
	return singleClient, nil
}

func newESDB(cfg config.IAppConfig) (*esdb.Client, error) {
	if cfg == nil {
		return nil, errors.ErrNoConfig
	}
	db, err := transientESClient(cfg.GetESDBConfig())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func oneESDB(cfg config.IAppConfig) (*esdb.Client, error) {
	if cfg == nil {
		return nil, errors.ErrNoConfig
	}
	db, err := singletonESClient(cfg.GetESDBConfig())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil

}

// TransientESClient returns a functor that creates a transient EventStoreDB client
func TransientESClient(cfg config.IAppConfig) EventStoreDBFtor {
	return func() *esdb.Client {
		c, err := newESDB(cfg)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		return c
	}
}

// SingletonESClient is an injection that returns an EventStoreDB functor
func SingletonESClient(cfg config.IAppConfig) EventStoreDBFtor {
	return func() *esdb.Client {
		c, err := oneESDB(cfg)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		return c
	}
}

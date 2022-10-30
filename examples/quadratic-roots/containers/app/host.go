package app

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/spokes"
	"sync"
)

var (
	singleton spokes.IApp
	cMutex    = &sync.Mutex{}
)

type host struct {
	*container.App
}

func Host(cfg config.IAppConfig) spokes.IApp {
	if singleton == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		a := &host{}
		b := container.NewApp(cfg, nil, nil)
		a.App = b
		singleton = a
	}
	return singleton
}

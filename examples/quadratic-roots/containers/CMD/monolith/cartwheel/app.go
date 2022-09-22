package cartwheel

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

type app struct {
	*container.App
}

func SingleApp(cfg config.IAppConfig) spokes.IApp {
	if singleton == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		a := &app{}
		b := container.NewApp(cfg, nil, nil)
		a.App = b
		singleton = a
	}
	return singleton
}

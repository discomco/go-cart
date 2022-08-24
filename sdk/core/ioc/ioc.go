package ioc

import (
	"go.uber.org/dig"
	"sync"
)

var (
	cDigMutex    = &sync.Mutex{}
	digSingleton IDig
)

// C is just an extension struct that allows us to extend dig functionality
type C struct {
	*dig.Container
}

// Inject allows us to inject IAggPlugins
func (c *C) Inject(dig IDig, items ...interface{}) IDig {
	for _, item := range items {
		_ = dig.Provide(item)
	}
	return dig
}

// NewIoC creates a new DI container
func NewIoC() IDig {
	return &C{
		Container: dig.New(),
	}
}

// SingleIoC creates a new singleton DI container
func SingleIoC() IDig {
	if digSingleton == nil {
		cDigMutex.Lock()
		defer cDigMutex.Unlock()
		digSingleton = NewIoC()
	}
	return digSingleton
}

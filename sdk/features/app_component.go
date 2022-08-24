package features

import (
	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/core/errors"
	"github.com/discomco/go-cart/core/ioc"
	"github.com/discomco/go-cart/core/logger"
	"github.com/discomco/go-cart/core/mediator"
)

type Name string

type AppComponent struct {
	Name     Name
	cfg      config.IAppConfig
	logger   logger.IAppLogger
	mediator mediator.IMediator
}

func (a *AppComponent) GetConfig() config.IAppConfig {
	if a.cfg == nil {
		panic(errors.ErrNoConfig)
	}
	return a.cfg
}

func (a *AppComponent) GetLogger() logger.IAppLogger {
	if a.logger == nil {
		panic(errors.ErrNoLogger)
	}
	return a.logger
}

func (a *AppComponent) GetMediator() mediator.IMediator {
	if a.mediator == nil {
		panic(errors.ErrNoMediator)
	}
	return a.mediator
}

func (a *AppComponent) GetName() Name {
	return a.Name
}

func NewAppComponent(name Name) *AppComponent {
	dig := ioc.SingleIoC()
	c := &AppComponent{
		Name: name,
	}
	_ = dig.Invoke(func(appConfig config.IAppConfig, appLogger logger.IAppLogger, mediator mediator.IMediator) {
		c.cfg = appConfig
		c.logger = appLogger
		c.mediator = mediator
	})
	return c
}

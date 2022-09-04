package reactors

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/errors"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/schema"
)

type Component struct {
	Name     schema.Name
	cfg      config.IAppConfig
	logger   logger.IAppLogger
	mediator mediator.IMediator
}

func (a *Component) GetConfig() config.IAppConfig {
	if a.cfg == nil {
		panic(errors.ErrNoConfig)
	}
	return a.cfg
}

func (a *Component) GetLogger() logger.IAppLogger {
	if a.logger == nil {
		panic(errors.ErrNoLogger)
	}
	return a.logger
}

func (a *Component) GetMediator() mediator.IMediator {
	if a.mediator == nil {
		panic(errors.ErrNoMediator)
	}
	return a.mediator
}

func (a *Component) GetName() schema.Name {
	return a.Name
}

func NewComponent(name schema.Name) *Component {
	dig := ioc.SingleIoC()
	c := &Component{
		Name: name,
	}
	_ = dig.Invoke(func(appConfig config.IAppConfig, appLogger logger.IAppLogger, mediator mediator.IMediator) {
		c.cfg = appConfig
		c.logger = appLogger
		c.mediator = mediator
	})
	return c
}

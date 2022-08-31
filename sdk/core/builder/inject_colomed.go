package builder

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
)

// InjectCoLoMed creates a basic DI Container that offers Configuration, Logging and a mediator
func InjectCoLoMed(cfgPath string) ioc.IDig {
	dig := ioc.SingleIoC()
	return dig.Inject(dig,
		func() config.Path { return config.Path(cfgPath) },
		config.AppConfig,
		logger.AppLogger,
		mediator.SingletonMediator,
	)
}

package cartwheel

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/containers/app"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/spokes"
	"log"
)

func Run(cfgPath string) error {
	dig := build(cfgPath)
	var runner spokes.IApp
	var appLogger logger.IAppLogger
	err := dig.Invoke(func(
		r spokes.IApp,
		l logger.IAppLogger) {
		runner = r
		appLogger = l
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	appLogger.Fatal(runner.Run())
	return nil
}

func build(path string) ioc.IDig {
	dig := container.DefaultQRY(path)
	dig.Inject(dig,
		app.Host,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ioc.IDig {
	return dig
}

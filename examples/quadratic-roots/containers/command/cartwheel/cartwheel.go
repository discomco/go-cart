package cartwheel

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/containers/app"
	calc_roots "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/spoke"
	initialize_calc "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/spoke"
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

func build(cfgPath string) ioc.IDig {
	dig := container.DefaultCMD(cfgPath)
	dig = dig.Inject(dig,
		app.Host,
	).Inject(dig,
		calc_roots.Spoke,
		initialize_calc.Spoke,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		_ initialize_calc.ISpoke,
		_ calc_roots.ISpoke,
	) {
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

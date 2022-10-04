package cartwheel

import (
	calc_roots "github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/spoke"
	initialize_calc "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/spoke"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"log"
)

func BuildCartwheel(cfgPath string) ioc.IDig {
	dig := container.DefaultCMD(cfgPath)
	dig = dig.Inject(dig,
		SingleApp,
	).Inject(dig,
		calc_roots.Spoke,
		initialize_calc.Spoke,
	)
	return resolveSpokes(dig)
}

func resolveSpokes(dig ioc.IDig) ioc.IDig {
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

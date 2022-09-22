package cartwheel

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/builder"
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/ftor"
	"github.com/discomco/go-cart/examples/quadratic-roots/drivers/redis"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	initialize_calc "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/spoke"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/eventstore_db"
	"log"
)

func BuildCartwheel(cfgPath string) ioc.IDig {
	dig := container.DefaultCMD(cfgPath)
	dig = dig.Inject(dig,
		SingleApp,
	).Inject(dig,
		eventstore_db.EvtProjFtor,
		eventstore_db.EventProjector,
	).Inject(dig,
		redis.DocStore,
		redis.ListStore,
	).Inject(dig,
		schema.DocFtor,
		schema.ListFtor,
	).Inject(dig,
		ftor.BehaviorFtor,
		builder.BehaviorBuilder,
	).Inject(dig,
		initialize_calc.Spoke,
	)
	return resolveSpokes(dig)
}

func resolveSpokes(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		_ initialize_calc.ISpoke,
	) {
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

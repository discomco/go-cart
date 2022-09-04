package cartwheel

import (
	"github.com/discomco/go-cart/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/drivers/redis"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	initialize_game "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/spoke"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/eventstore_db"
	"log"
)

func BuildCartWheel(cfgPath string) ioc.IDig {
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
		initialize_game.Spoke,
	)
	return resolveSpokes(dig)
}

func resolveSpokes(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		_ initialize_game.ISpoke) {
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

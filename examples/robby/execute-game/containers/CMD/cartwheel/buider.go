package cartwheel

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/examples/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/examples/robby/execute-game/drivers/redis"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	change_game_details "github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/spoke"
	change_game_settings "github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_settings/spoke"
	initialize_game "github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/spoke"
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
		change_game_details.Spoke,
		change_game_settings.Spoke,
	)
	return resolveSpokes(dig)
}

func resolveSpokes(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		_ initialize_game.ISpoke,
		_ change_game_details.ISpoke,
		_ change_game_settings.ISpoke,
	) {
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

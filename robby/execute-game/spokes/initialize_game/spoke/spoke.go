package spoke

import (
	"github.com/discomco/go-cart/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/drivers/redis"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/reactors"
	sdk_reactors "github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/eventstore_db"
	"github.com/discomco/go-cart/sdk/features"
	"log"
)

type ISpoke interface {
	features.ICmdSpoke
}

func newCmdSpoke() ISpoke {
	return features.NewCmdFeature(behavior.CMD_TOPIC)
}

func Spoke(cfgPath config.Path) ISpoke {
	dig := container.DefaultCMD(string(cfgPath))
	dig.Inject(dig,
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
		behavior.EvtToDoc,
		reactors.ToRedisDoc,
	).Inject(dig,
		reactors.Responder,
		behavior.Hope2Cmd,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ISpoke {
	spoke := newCmdSpoke()
	err := dig.Invoke(func(
		responder reactors.IResponder,
		projector sdk_reactors.IProjector,
		toRedisDoc reactors.IToRedisDoc,
	) {
		spoke.Inject(
			projector,
			responder,
			toRedisDoc,
		)
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return spoke
}

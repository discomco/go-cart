package spoke

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/builder"
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/ftor"
	"github.com/discomco/go-cart/examples/quadratic-roots/drivers/redis"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/comps"
	sdk_comps "github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/eventstore_db"
	"github.com/discomco/go-cart/sdk/spokes"
	"log"
)

type ISpoke interface {
	spokes.ICommandSpoke
}

func newSpoke() ISpoke {
	return spokes.NewCmdSpoke(behavior.CmdTopic)
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
		comps.ToRedisDoc,
		behavior.EvtToList,
		comps.ToRedisList,
	).Inject(dig,
		comps.Responder,
		behavior.Hope2Cmd,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ISpoke {
	spoke := newSpoke()
	err := dig.Invoke(func(
		projector sdk_comps.IProjector,
		responder comps.IResponder,
		toRedisDoc comps.IToRedisDoc,
		toRedisList comps.IToRedisList,
	) {
		spoke.Inject(
			projector,
			responder,
			toRedisDoc,
			toRedisList,
		)
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return spoke
}

package comps

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/ftor"
	"github.com/discomco/go-cart/examples/quadratic-roots/drivers/redis"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/behavior"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"log"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv            ioc.IDig
	testLogger         logger.IAppLogger
	newTestCalculation sdk_behavior.BehaviorBuilder
	newTestCH          comps.CmdHandlerFtor
	newTestRequester   comps.GenRequesterFtor[contract.IHope]
)

func init() {
	testEnv = buildTestEnv()
}

func buildLocalBehavior(ftor sdk_behavior.GenBehaviorFtor[schema.QuadraticDoc]) sdk_behavior.BehaviorBuilder {
	return func() sdk_behavior.IBehavior {
		agg := ftor()
		agg.Inject(agg,
			behavior.TryCmd,
			behavior.ApplyEvt,
		)
		return agg
	}

}

func buildTestEnv() ioc.IDig {
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		schema.DocFtor,
		schema.ListFtor,
		ftor.BehaviorFtor,
		buildLocalBehavior,
	).Inject(dig,
		behavior.EvtToDoc,
		behavior.EvtToList,
		redis.DocStore,
		redis.ListStore,
		ToRedisDoc,
		ToRedisList,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	if err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newBeh sdk_behavior.BehaviorBuilder,
		newCH comps.CmdHandlerFtor,
	) {
		testLogger = appLogger
		newTestCalculation = newBeh
		newTestCH = newCH
	}); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

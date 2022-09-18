package behavior

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	initialize_game_behavior "github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"log"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv         ioc.IDig
	testLogger      logger.IAppLogger
	newTestBehavior sdk_behavior.BehaviorBuilder
)

func init() {
	testEnv = buildTestEnv()
}

func localBuilder(ftor sdk_behavior.GenBehaviorFtor[schema.GameDoc]) sdk_behavior.BehaviorBuilder {
	return func() sdk_behavior.IBehavior {
		agg := ftor()
		agg.Inject(agg,
			initialize_game_behavior.TryCmd,
			initialize_game_behavior.ApplyEvt,
			TryCmd,
			ApplyEvt,
		)
		return agg
	}

}

func buildTestEnv() ioc.IDig {
	dig := builder.InjectCoLoMed(ConfigPath)
	dig.Inject(dig,
		schema.DocFtor,
		ftor.BehaviorFtor,
		localBuilder,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newBeh sdk_behavior.BehaviorBuilder,
	) {
		testLogger = appLogger
		newTestBehavior = newBeh
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

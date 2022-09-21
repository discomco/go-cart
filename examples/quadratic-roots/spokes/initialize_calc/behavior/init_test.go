package behavior

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/behavior/ftor"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
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
	testEnv            ioc.IDig
	testLogger         logger.IAppLogger
	newTestCalculation sdk_behavior.BehaviorBuilder
)

func init() {
	testEnv = buildTestEnv()
}

func localBuilder(ftor sdk_behavior.GenBehaviorFtor[schema.QuadraticDoc]) sdk_behavior.BehaviorBuilder {
	return func() sdk_behavior.IBehavior {
		agg := ftor()
		agg.Inject(agg,
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
		ftor.CalculationFtor,
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
		newTestCalculation = newBeh
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

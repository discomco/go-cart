package behavior

import (
	"github.com/discomco/go-cart/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	domain2 "github.com/discomco/go-cart/sdk/domain"
	"log"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv    ioc.IDig
	testLogger logger.IAppLogger
	newTestAgg domain2.AggBuilder
)

func init() {
	testEnv = buildTestEnv()
}

func localBuilder(ftor domain2.GenAggFtor[schema.GameDoc]) domain2.AggBuilder {
	return func() domain2.IAggregate {
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
		schema.GameDocFtor,
		ftor.BehaviorFtor,
		localBuilder,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newAgg domain2.AggBuilder,
	) {
		testLogger = appLogger
		newTestAgg = newAgg
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

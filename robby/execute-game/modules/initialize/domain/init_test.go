package domain

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/domain/ftor"
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	domain2 "github.com/discomco/go-cart/sdk/domain"
	"log"
)

const (
	ConfigPath = "../../../-shared/config/config.yaml"
)

var (
	testEnv    ioc.IDig
	testLogger logger.IAppLogger
	newTestAgg domain2.AggBuilder
)

func init() {
	testEnv = buildTestEnv()
}

func LocalBuilder(ftor domain2.GenAggFtor[model.Root]) domain2.AggBuilder {
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
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		model.RootFtor,
		ftor.AggFtor,
		LocalBuilder,
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

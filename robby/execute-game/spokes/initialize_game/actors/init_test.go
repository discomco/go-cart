package actors

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/-shared/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/-shared/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/features"
	"log"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv           ioc.IDig
	testLogger        logger.IAppLogger
	newTestBehavior   domain.AggBuilder
	newTestCmdHandler features.CmdHandlerFtor
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		schema.RootFtor,
	).Inject(dig,
		ftor.BehaviorFtor,
		builder.BehaviorBuilder,
	).Inject(dig,
		Responder,
		behavior.Hope2Cmd,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newBehavior domain.AggBuilder,
		newCmdHandler features.CmdHandlerFtor,
	) {
		testLogger = appLogger
		newTestBehavior = newBehavior
		newTestCmdHandler = newCmdHandler
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

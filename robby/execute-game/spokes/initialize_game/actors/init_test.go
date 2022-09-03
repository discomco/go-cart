package actors

import (
	"github.com/discomco/go-cart/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/schema"
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
	testRequester     IRequester
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		schema.GameDocFtor,
	).Inject(dig,
		ftor.BehaviorFtor,
		builder.BehaviorBuilder,
	).Inject(dig,
		Responder,
		behavior.Hope2Cmd,
		Requester,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newBehavior domain.AggBuilder,
		newCmdHandler features.CmdHandlerFtor,
		requester IRequester,
	) {
		testLogger = appLogger
		newTestBehavior = newBehavior
		newTestCmdHandler = newCmdHandler
		testRequester = requester
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}
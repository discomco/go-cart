package reactors

import (
	"github.com/discomco/go-cart/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/reactors"
	"log"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv           ioc.IDig
	testLogger        logger.IAppLogger
	newTestBehavior   sdk_behavior.BehaviorBuilder
	newTestCmdHandler reactors.CmdHandlerFtor
	testRequester     IRequester
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		schema.DocFtor,
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
		newBehavior sdk_behavior.BehaviorBuilder,
		newCmdHandler reactors.CmdHandlerFtor,
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

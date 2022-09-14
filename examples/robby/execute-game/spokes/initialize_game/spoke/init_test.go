package spoke

import (
	"github.com/discomco/go-cart/robby/execute-game/behavior/builder"
	"github.com/discomco/go-cart/robby/execute-game/behavior/ftor"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/comps"
	sdk_behavior "github.com/discomco/go-cart/sdk/behavior"
	sdk_reactors "github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"log"
)

const (
	ConfigPath = "../../../config/config.yaml"
)

var (
	testEnv           ioc.IDig
	testLogger        logger.IAppLogger
	newTestBehavior   sdk_behavior.BehaviorBuilder
	newTestCmdHandler sdk_reactors.CmdHandlerFtor
	testModule        ISpoke
	testRequester     comps.IRequester
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
		comps.Responder,
		behavior.Hope2Cmd,
		Spoke,
		comps.Requester,
	)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(
		appLogger logger.IAppLogger,
		newBehavior sdk_behavior.BehaviorBuilder,
		newCmdHandler sdk_reactors.CmdHandlerFtor,
		mod ISpoke,
		req comps.IRequester,
	) {
		testLogger = appLogger
		newTestBehavior = newBehavior
		newTestCmdHandler = newCmdHandler
		testModule = mod
		testRequester = req
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

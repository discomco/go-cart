package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
)

const (
	CfgPath = "../../config/config.yaml"
)

var (
	testEnv            ioc.IDig
	testProjector      comps.IProjector
	testLogger         logger.IAppLogger
	testAS             comps.IBehaviorStore
	testES             comps.IEventStore
	testConfig         config.IAppConfig
	testMed            mediator.IMediator
	testStreamIDs      []string
	testLoggingHandler comps.ILoggingReactor
)

func init() {
	testEnv = buildTestEnv()
	testStreamIDs = []string{"sdk-stream-1", "sdk-stream-2"}
	err := testEnv.Invoke(func(
		log logger.IAppLogger,
		cfg config.IAppConfig,
		asFtor comps.BehSFtor,
		esFtor comps.ESFtor,
		prjFtor comps.ProjectorFtor,
		med mediator.IMediator,
		medLogger comps.ILoggingReactor) {
		testConfig = cfg
		testLogger = log
		testProjector = prjFtor()
		testAS = asFtor()
		testES = esFtor()
		testMed = med
		testLoggingHandler = medLogger
	})
	if err != nil {
		testLogger.Error(err)
	}
}

func buildTestEnv() ioc.IDig {
	dig := EventSourcing(CfgPath)
	return dig.Inject(
		AddESProjector(dig),
		mediator.SingletonMediator,
	).Inject(dig,
		comps.GeneralMediatorLogger,
	)
}

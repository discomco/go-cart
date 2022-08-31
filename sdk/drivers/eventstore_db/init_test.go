package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/features"
)

const (
	CfgPath = "../../config/config.yaml"
)

var (
	testEnv            ioc.IDig
	testProjector      features.IEventProjector
	testLogger         logger.IAppLogger
	testAS             features.IAggregateStore
	testES             features.IEventStore
	testConfig         config.IAppConfig
	testMed            mediator.IMediator
	testStreamIDs      []string
	testLoggingHandler features.IMediatorLogger
)

func init() {
	testEnv = buildTestEnv()
	testStreamIDs = []string{"sdk-stream-1", "sdk-stream-2"}
	err := testEnv.Invoke(func(
		log logger.IAppLogger,
		cfg config.IAppConfig,
		asFtor features.ASFtor,
		esFtor features.ESFtor,
		prjFtor features.EventProjectorFtor,
		med mediator.IMediator,
		medLogger features.IMediatorLogger) {
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
		features.GeneralMediatorLogger,
	)
}

package nats

import (
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
)

const (
	CfgPath    = "../../config/config.yaml"
	TEST_TOPIC = "test.topic"
)

var (
	testEnv    ioc.IDig
	testLogger logger.IAppLogger
)

func buildTestEnv() ioc.IDig {
	ioc := builder.InjectCoLoMed(CfgPath)
	return ioc.Inject(ioc,
		SingleNATS,
	)
	return ioc
}

func init() {
	testEnv = buildTestEnv()
	err := testEnv.Invoke(func(appLogger logger.IAppLogger) {
		testLogger = appLogger
	})
	if err != nil {
		panic(err)
	}
}

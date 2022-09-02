package avatar

import (
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"log"
)

const (
	ConfigPath = "../../config/config.yaml"
)

var (
	testEnv    ioc.IDig
	testLogger logger.IAppLogger
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := builder.InjectCoLoMed(ConfigPath)
	return resolveTestEnv(dig)
}

func resolveTestEnv(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(appLogger logger.IAppLogger) {
		testLogger = appLogger
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

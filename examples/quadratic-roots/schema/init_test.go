package schema

import (
	"github.com/discomco/go-cart/sdk/container"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/logger"
	"log"
)

const (
	ConfigPath = "../config/config.yaml"
)

var (
	testEnv    ioc.IDig
	testLogger logger.IAppLogger
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := container.DefaultCMD(ConfigPath)
	dig.Inject(dig,
		DocFtor,
		ListFtor,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(appLogger logger.IAppLogger) {
		testLogger = appLogger
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

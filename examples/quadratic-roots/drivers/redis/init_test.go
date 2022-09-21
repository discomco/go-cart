package redis

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
	dig.Inject(dig,
		DocStore,
		ListStore,
	)
	return resolve(dig)
}

func resolve(dig ioc.IDig) ioc.IDig {
	err := dig.Invoke(func(tl logger.IAppLogger) {
		testLogger = tl
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return dig
}

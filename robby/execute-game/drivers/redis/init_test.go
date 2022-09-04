package redis

import (
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
)

const (
	ConfigPath = "../../config.yaml"
)

var (
	testEnv ioc.IDig
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
	return dig
}

package kafka

import (
	"github.com/discomco/go-cart/core/builder"
	"github.com/discomco/go-cart/core/ioc"
)

const (
	ConfigPath = "../../config/config.yaml"
)

var (
	testEnv ioc.IDig
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := builder.InjectCoLoMed(ConfigPath)
	return dig
}

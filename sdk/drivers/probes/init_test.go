package probes

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
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
	return dig.Inject(dig,
		BogusMetrics,
	)
}

type IBogusMetrics interface {
	IMetrics
}

func BogusMetrics(cfg config.IAppConfig) IBogusMetrics {
	return NewMetrics(cfg)
}

type IBogusCounter interface {
	IMetricsCounter
}

func BogusCounter(config config.IAppConfig) IBogusCounter {
	sCfg := config.GetServiceConfig()
	return NewCounter("bogusCounter", sCfg.GetNamespace(), sCfg.GetSubSystem(), "Help", nil)
}

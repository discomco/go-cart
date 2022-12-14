package container

import (
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
)

const (
	ConfigPath = "../config/config.yaml"
)

var (
	testEnv ioc.IDig
)

func init() {
	testEnv = builder.InjectCoLoMed(ConfigPath)
}

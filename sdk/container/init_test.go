package container

import (
	"github.com/discomco/go-cart/core/builder"
	"github.com/discomco/go-cart/core/ioc"
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

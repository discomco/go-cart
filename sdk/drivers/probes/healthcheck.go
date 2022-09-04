package probes

import (
	"github.com/discomco/go-cart/sdk/features"
	"github.com/discomco/go-cart/sdk/reactors"
	"github.com/discomco/go-cart/sdk/schema"
)

type IHealthcheck interface {
	features.ISpoke
}

type Healthcheck struct {
	*reactors.Component
}

const (
	HealthCheckNameFmt = "[%+v].Healthcheck"
)

func newHealthCheck(name schema.Name) *Healthcheck {
	hc := &Healthcheck{}
	cp := reactors.NewComponent(name)
	hc.Component = cp
	return hc
}

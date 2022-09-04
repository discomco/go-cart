package probes

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-cart/sdk/spokes"
)

type IHealthcheck interface {
	spokes.ISpoke
}

type Healthcheck struct {
	*comps.Component
}

const (
	HealthCheckNameFmt = "[%+v].Healthcheck"
)

func newHealthCheck(name schema.Name) *Healthcheck {
	hc := &Healthcheck{}
	cp := comps.NewComponent(name)
	hc.Component = cp
	return hc
}

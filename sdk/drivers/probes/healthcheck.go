package probes

import "github.com/discomco/go-cart/features"

type IHealthcheck interface {
	features.IFeature
}

type Healthcheck struct {
	*features.AppComponent
}

const (
	HealthCheckNameFmt = "[%+v].Healthcheck"
)

func newHealthCheck(name features.Name) *Healthcheck {
	hc := &Healthcheck{}
	cp := features.NewAppComponent(name)
	hc.AppComponent = cp
	return hc
}

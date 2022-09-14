package spokes

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/comps/pulsar"
	"github.com/discomco/go-cart/sdk/schema"
	"time"
)

type IPulsarSpoke interface {
	ISpoke
}

//PulsarSpoke is a spoke that contains a number of PulseWorkers
type PulsarSpoke struct {
	*Spoke
	ticker  *time.Ticker
	pulsars map[string]pulsar.IPulsar
}

func (ps *PulsarSpoke) startSpoke(ctx context.Context) error {
	panic("not implemented")
}

func (ps *PulsarSpoke) stopSpoke(ctx context.Context) {
	panic("not implemented")
}

func (ps *PulsarSpoke) registerPlugins(plugins []comps.ISpokePlugin) {

}

func newPulsarSpoke(
	name schema.Name,
	interval time.Duration,
) *PulsarSpoke {
	ps := &PulsarSpoke{
		ticker:  time.NewTicker(interval),
		pulsars: make(map[string]pulsar.IPulsar),
	}
	b := NewSpoke(name, ps.startSpoke, ps.stopSpoke, ps.registerPlugins)
	ps.Spoke = b
	return ps
}

package spokes

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/tirol"
	"github.com/discomco/go-cart/sdk/schema"
)

type QrySpoke struct {
	*Spoke
	providers map[schema.Name]comps.IQueryProvider
	tirol     tirol.ITirol
}

func (qs *QrySpoke) registerPlugins(plugins []comps.ISpokePlugin) {
	for _, plugin := range plugins {
		switch plugin.(type) {
		case comps.IQueryProvider:
			qs.registerQueryProvider(plugin.(comps.IQueryProvider))
		default:
			continue
		}
	}
}

func (qs *QrySpoke) registerQueryProvider(provider comps.IQueryProvider) {
	if qs.providers[provider.GetName()] == nil {
		qs.providers[provider.GetName()] = provider
	}
}

func (qs *QrySpoke) runLocal(ctx context.Context) error {
	for _, provider := range qs.providers {
		err := provider.Activate(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (qs *QrySpoke) downLocal(ctx context.Context) {
	for _, provider := range qs.providers {
		_ = provider.Deactivate(ctx)
	}
}

func newQrySpoke(name schema.Name, path string) (*QrySpoke, error) {
	f := &QrySpoke{}
	b := NewSpoke(name, f.runLocal, f.downLocal, f.registerPlugins)
	dig := ioc.SingleIoC()
	err := dig.Invoke(func(tirol tirol.ITirol) {
		f.tirol = tirol
	})
	if err != nil {
		return nil, err
	}
	f.Spoke = b
	return f, nil
}

func NewQrySpoke(name schema.Name, path string) (IQuerySpoke, error) {
	return newQrySpoke(name, path)
}

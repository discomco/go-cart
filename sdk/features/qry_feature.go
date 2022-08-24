package features

import (
	"context"

	"github.com/discomco/go-cart/core/ioc"
	"github.com/discomco/go-cart/drivers/tirol"
)

type QryFeature struct {
	*Feature
	providers map[Name]IQueryProvider
	tirol     tirol.ITirol
}

func (f *QryFeature) registerFeaturePlugins(plugins []IFeaturePlugin) {
	for _, plugin := range plugins {
		switch plugin.(type) {
		case IQueryProvider:
			f.registerQueryProvider(plugin.(IQueryProvider))
		default:
			continue
		}
	}
}

func (f *QryFeature) registerQueryProvider(provider IQueryProvider) {
	if f.providers[provider.GetName()] == nil {
		f.providers[provider.GetName()] = provider
	}
}

func (f *QryFeature) runLocal(ctx context.Context) error {
	for _, provider := range f.providers {
		err := provider.Activate(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *QryFeature) downLocal(ctx context.Context) {
	for _, provider := range f.providers {
		_ = provider.Deactivate(ctx)
	}
}

func newQryFeature(name Name, path string) (*QryFeature, error) {
	f := &QryFeature{}
	b := NewFeature(name, f.runLocal, f.downLocal, f.registerFeaturePlugins)
	dig := ioc.SingleIoC()
	err := dig.Invoke(func(tirol tirol.ITirol) {
		f.tirol = tirol
	})
	if err != nil {
		return nil, err
	}
	f.Feature = b
	return f, nil
}

func NewQryAppFeature(name Name, path string) (IQryFeature, error) {
	return newQryFeature(name, path)
}

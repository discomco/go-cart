package features

import (
	"context"

	"github.com/discomco/go-cart/core/ioc"
)

type (
	runFeatureFunc      func(ctx context.Context) error
	downFeatureFunc     func(ctx context.Context)
	registerPluginsFunc func(plugins []IFeaturePlugin)
	FeatureFtor         func() IFeature
	FeatureBuilder      func() IFeature
)

type Feature struct {
	*AppComponent
	run        runFeatureFunc
	down       downFeatureFunc
	regPlugins registerPluginsFunc
}

func (f *Feature) Shutdown(ctx context.Context) {
	if f.down == nil {
		return
	}
	f.GetLogger().Infof("Gracefully shutting down Feature [%+v]", f.GetName())
	f.down(ctx)
	f.GetLogger().Infof("Feature [%+v] is DOWN!", f.GetName())
}

func (f *Feature) Run(ctx context.Context) func() error {
	return func() error {
		if f.run == nil {
			return nil
		}
		f.GetLogger().Infof("Feature [%+v] is UP!", f.GetName())
		return f.run(ctx)
	}
}

func (f *Feature) Inject(plugins ...IFeaturePlugin) IFeature {
	if len(plugins) == 0 {
		return f
	}
	if f.regPlugins != nil {
		f.regPlugins(plugins)
	}
	return f
}

func NewFeature(
	name Name,
	run runFeatureFunc,
	down downFeatureFunc,
	regPlugins registerPluginsFunc,
) *Feature {
	if name == "" {
		name = "Feature"
	}
	base := NewAppComponent(name)
	base.Name = name
	f := &Feature{
		run:        run,
		down:       down,
		regPlugins: regPlugins,
	}
	f.AppComponent = base

	dig := ioc.SingleIoC()
	_ = dig.Invoke(func(app IApp) {
		if app != nil {
			app.Inject(f)
		}
	})

	return f
}

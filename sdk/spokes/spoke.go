package spokes

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/schema"
)

type (
	runSpokeFunc        func(ctx context.Context) error
	downSpokeFunc       func(ctx context.Context)
	registerPluginsFunc func(plugins []comps.ISpokePlugin)
	SpokeFtor           func() ISpoke
	SpokeBuilder        func() ISpoke
)

// Spoke is the base struct to use when implementing new types of modules.
type Spoke struct {
	*comps.Component
	run        runSpokeFunc
	down       downSpokeFunc
	regPlugins registerPluginsFunc
}

// Shutdown gracefully shuts down the Spoke
func (f *Spoke) Shutdown(ctx context.Context) {
	if f.down == nil {
		return
	}
	f.GetLogger().Infof("Gracefully shutting down Spoke [%+v]", f.GetName())
	f.down(ctx)
	f.GetLogger().Infof("Spoke [%+v] is DOWN!", f.GetName())
}

// Run is the go routine that executes a Spoke
func (f *Spoke) Run(ctx context.Context) func() error {
	return func() error {
		if f.run == nil {
			return nil
		}
		f.GetLogger().Infof("Spoke [%+v] is UP!", f.GetName())
		return f.run(ctx)
	}
}

// Inject allows you to inject Reactions into the Spoke
func (f *Spoke) Inject(plugins ...comps.ISpokePlugin) ISpoke {
	if len(plugins) == 0 {
		return f
	}
	if f.regPlugins != nil {
		f.regPlugins(plugins)
	}
	return f
}

// NewSpoke returns a new Spoke
func NewSpoke(
	name schema.Name,
	run runSpokeFunc,
	down downSpokeFunc,
	regPlugins registerPluginsFunc,
) *Spoke {
	if name == "" {
		name = "Spoke"
	}
	base := comps.NewComponent(name)
	base.Name = name
	f := &Spoke{
		run:        run,
		down:       down,
		regPlugins: regPlugins,
	}
	f.Component = base
	dig := ioc.SingleIoC()
	_ = dig.Invoke(func(app IApp) {
		if app != nil {
			app.Inject(f)
		}
	})
	return f
}

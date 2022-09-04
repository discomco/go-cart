package features

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/schema"
)

type (
	runSpokeFunc         func(ctx context.Context) error
	downSpokeFunc        func(ctx context.Context)
	registerReactorsFunc func(plugins []comps.IReactor)
	SpokeFtor            func() ISpoke
	SpokeBuilder         func() ISpoke
)

type Spoke struct {
	*comps.Component
	run         runSpokeFunc
	down        downSpokeFunc
	regReactors registerReactorsFunc
}

func (f *Spoke) Shutdown(ctx context.Context) {
	if f.down == nil {
		return
	}
	f.GetLogger().Infof("Gracefully shutting down Spoke [%+v]", f.GetName())
	f.down(ctx)
	f.GetLogger().Infof("Spoke [%+v] is DOWN!", f.GetName())
}

func (f *Spoke) Run(ctx context.Context) func() error {
	return func() error {
		if f.run == nil {
			return nil
		}
		f.GetLogger().Infof("Spoke [%+v] is UP!", f.GetName())
		return f.run(ctx)
	}
}

func (f *Spoke) Inject(reactors ...comps.IReactor) ISpoke {
	if len(reactors) == 0 {
		return f
	}
	if f.regReactors != nil {
		f.regReactors(reactors)
	}
	return f
}

func NewSpoke(
	name schema.Name,
	run runSpokeFunc,
	down downSpokeFunc,
	regReactors registerReactorsFunc,
) *Spoke {
	if name == "" {
		name = "Spoke"
	}
	base := comps.NewComponent(name)
	base.Name = name
	f := &Spoke{
		run:         run,
		down:        down,
		regReactors: regReactors,
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

package features

import (
	"golang.org/x/net/context"
	"time"
)

type (
	ProjFeatureFtor    func() IPrjFeature
	ProjFeatureBuilder func() IPrjFeature
)

type projFeature struct {
	*Feature
	projector   IEventProjector
	projections map[Name]IMediatorSubscriber
}

func (f *projFeature) run(ctx context.Context) error {
	for _, handler := range f.projections {
		err := handler.Activate(ctx)
		if err != nil {
			return err
		}
	}
	time.Sleep(5 * time.Second)
	return f.projector.Activate(ctx)
}

func (f *projFeature) down(ctx context.Context) {
	_ = f.projector.Deactivate(ctx)
	for _, handler := range f.projections {
		_ = handler.Deactivate(ctx)
	}
}

func (f *projFeature) registerProjection(handler IMediatorSubscriber) {
	_, ok := f.projections[handler.GetName()]
	if !ok {
		f.projections[handler.GetName()] = handler
	}
}

func (f *projFeature) registerProjector(projector IEventProjector) {
	f.projector = projector
}

func (f *projFeature) registerFeaturePlugins(plugins []IFeaturePlugin) {
	for _, plugin := range plugins {
		switch plugin.(type) {
		case IMediatorSubscriber:
			f.registerProjection(plugin.(IMediatorSubscriber))
		case IEventProjector:
			f.registerProjector(plugin.(IEventProjector))
		default:
			continue
		}
	}
}

func NewProjectionAppFeature(
	name Name,
) *projFeature {
	f := &projFeature{
		projections: make(map[Name]IMediatorSubscriber),
	}
	f.Feature = NewFeature(name, f.run, f.down, f.registerFeaturePlugins)
	return f
}

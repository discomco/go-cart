package features

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/net/context"
	"time"
)

type (
	ProjSpokeFtor    func() IPrjSpoke
	ProjSpokeBuilder func() IPrjSpoke
)

type projSpoke struct {
	*Spoke
	projector   comps.IProjector
	projections map[schema.Name]comps.IMediatorReactor
}

func (f *projSpoke) run(ctx context.Context) error {
	for _, handler := range f.projections {
		err := handler.Activate(ctx)
		if err != nil {
			return err
		}
	}
	time.Sleep(5 * time.Second)
	return f.projector.Activate(ctx)
}

func (f *projSpoke) down(ctx context.Context) {
	_ = f.projector.Deactivate(ctx)
	for _, handler := range f.projections {
		_ = handler.Deactivate(ctx)
	}
}

func (f *projSpoke) registerProjection(handler comps.IMediatorReactor) {
	_, ok := f.projections[handler.GetName()]
	if !ok {
		f.projections[handler.GetName()] = handler
	}
}

func (f *projSpoke) registerProjector(projector comps.IProjector) {
	f.projector = projector
}

func (f *projSpoke) registerReactors(plugins []comps.IReactor) {
	for _, plugin := range plugins {
		switch plugin.(type) {
		case comps.IMediatorReactor:
			f.registerProjection(plugin.(comps.IMediatorReactor))
		case comps.IProjector:
			f.registerProjector(plugin.(comps.IProjector))
		default:
			continue
		}
	}
}

func NewPrjSpoke(
	name schema.Name,
) *projSpoke {
	f := &projSpoke{
		projections: make(map[schema.Name]comps.IMediatorReactor),
	}
	f.Spoke = NewSpoke(name, f.run, f.down, f.registerReactors)
	return f
}

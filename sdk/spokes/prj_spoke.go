package spokes

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/net/context"
	"time"
)

type (
	ProjectionSpokeFtor    func() IProjectionSpoke
	ProjectionSpokeBuilder func() IProjectionSpoke
)

type ProjectionSpoke struct {
	*Spoke
	projector   comps.IProjector
	projections map[schema.Name]comps.IMediatorReaction
}

func (f *ProjectionSpoke) run(ctx context.Context) error {
	for _, handler := range f.projections {
		err := handler.Activate(ctx)
		if err != nil {
			return err
		}
	}
	time.Sleep(5 * time.Second)
	return f.projector.Activate(ctx)
}

func (f *ProjectionSpoke) down(ctx context.Context) {
	_ = f.projector.Deactivate(ctx)
	for _, handler := range f.projections {
		_ = handler.Deactivate(ctx)
	}
}

func (f *ProjectionSpoke) registerProjection(handler comps.IMediatorReaction) {
	_, ok := f.projections[handler.GetName()]
	if !ok {
		f.projections[handler.GetName()] = handler
	}
}

func (f *ProjectionSpoke) registerProjector(projector comps.IProjector) {
	f.projector = projector
}

func (f *ProjectionSpoke) registerReactions(reactions []comps.ISpokePlugin) {
	for _, plugin := range reactions {
		switch plugin.(type) {
		case comps.IMediatorReaction:
			f.registerProjection(plugin.(comps.IMediatorReaction))
		case comps.IProjector:
			f.registerProjector(plugin.(comps.IProjector))
		default:
			continue
		}
	}
}

func NewPrjSpoke(
	name schema.Name,
) *ProjectionSpoke {
	f := &ProjectionSpoke{
		projections: make(map[schema.Name]comps.IMediatorReaction),
	}
	f.Spoke = NewSpoke(name, f.run, f.down, f.registerReactions)
	return f
}

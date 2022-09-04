package features

import (
	"github.com/discomco/go-cart/sdk/reactors"
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
	projector   reactors.IProjector
	projections map[schema.Name]reactors.IMediatorReactor
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

func (f *projSpoke) registerProjection(handler reactors.IMediatorReactor) {
	_, ok := f.projections[handler.GetName()]
	if !ok {
		f.projections[handler.GetName()] = handler
	}
}

func (f *projSpoke) registerProjector(projector reactors.IProjector) {
	f.projector = projector
}

func (f *projSpoke) registerReactors(plugins []reactors.IReactor) {
	for _, plugin := range plugins {
		switch plugin.(type) {
		case reactors.IMediatorReactor:
			f.registerProjection(plugin.(reactors.IMediatorReactor))
		case reactors.IProjector:
			f.registerProjector(plugin.(reactors.IProjector))
		default:
			continue
		}
	}
}

func NewPrjSpoke(
	name schema.Name,
) *projSpoke {
	f := &projSpoke{
		projections: make(map[schema.Name]reactors.IMediatorReactor),
	}
	f.Spoke = NewSpoke(name, f.run, f.down, f.registerReactors)
	return f
}

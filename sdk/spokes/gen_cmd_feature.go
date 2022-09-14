package spokes

import (
	"context"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/hashicorp/go-multierror"
)

type (
	GenCmdSpokeBuilder[T behavior.ICmd] func() IGenCmdSpoke[T]
	GenCmdSpokeFtor[T behavior.ICmd]    func() IGenCmdSpoke[T]
)

type GenCmdSpoke struct {
	*Spoke
	responders []comps.IResponder
	listeners  []comps.IListener
	reactions  []comps.IMediatorReaction
}

func (f *GenCmdSpoke) up(ctx context.Context) error {
	errors := multierror.Error{}
	for _, handler := range f.reactions {
		err := handler.Activate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	for _, listener := range f.listeners {
		err := listener.Activate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	for _, responder := range f.responders {
		err := responder.Activate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	if len(errors.Errors) > 0 {
		err := multierror.Flatten(&errors)
		return err
	}
	return nil
}

func (f *GenCmdSpoke) down(ctx context.Context) {
	errors := multierror.Error{}
	for _, handler := range f.reactions {
		err := handler.Deactivate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	for _, listener := range f.listeners {
		err := listener.Deactivate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	for _, responder := range f.responders {
		err := responder.Deactivate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	if len(errors.Errors) > 0 {
		err := multierror.Flatten(&errors)
		f.GetLogger().Error(err)
	}
}

func (f *GenCmdSpoke) registerReactions(reactions []comps.ISpokePlugin) {
	if len(reactions) == 0 {
		return
	}
	for _, plugin := range reactions {
		switch plugin.(type) {
		case comps.IResponder:
			f.responders = append(f.responders, plugin.(comps.IResponder))
		case comps.IListener:
			f.listeners = append(f.listeners, plugin.(comps.IListener))
		case comps.IMediatorReaction:
			f.reactions = append(f.reactions, plugin.(comps.IMediatorReaction))
		}
	}
}

func NewGenCmdSpoke(name schema.Name) *GenCmdSpoke {
	f := &GenCmdSpoke{
		reactions:  make([]comps.IMediatorReaction, 0),
		responders: make([]comps.IResponder, 0),
		listeners:  make([]comps.IListener, 0),
	}
	base := NewSpoke(name, f.up, f.down, f.registerReactions)
	f.Spoke = base
	return f
}

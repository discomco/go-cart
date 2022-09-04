package features

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/hashicorp/go-multierror"
)

type (
	CmdSpokeFtor    func() ICmdSpoke
	CmdSpokeBuilder func() ICmdSpoke
)

type CmdSpoke struct {
	*Spoke
	projector  comps.IProjector
	responders []comps.IResponder
	listeners  []comps.IListener
	handlers   []comps.IMediatorReactor
}

func (f *CmdSpoke) up(ctx context.Context) error {
	errors := multierror.Error{}
	for _, handler := range f.handlers {
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
	if f.projector != nil {
		err := f.projector.Activate(ctx)
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

func (f *CmdSpoke) down(ctx context.Context) {
	errors := multierror.Error{}
	for _, handler := range f.handlers {
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
	if f.projector != nil {
		err := f.projector.Deactivate(ctx)
		if err != nil {
			errors.Errors = append(errors.Errors, err)
		}
	}
	if len(errors.Errors) > 0 {
		err := multierror.Flatten(&errors)
		f.GetLogger().Error(err)
	}
}

func (f *CmdSpoke) registerReactors(plugins []comps.IReactor) {
	if len(plugins) == 0 {
		return
	}
	for _, plugin := range plugins {
		switch plugin.(type) {
		case comps.IMediatorReactor, comps.IEvtReactor:
			f.handlers = append(f.handlers, plugin.(comps.IMediatorReactor))
		case comps.IResponder:
			f.responders = append(f.responders, plugin.(comps.IResponder))
		case comps.IListener:
			f.listeners = append(f.listeners, plugin.(comps.IListener))
		case comps.IProjector:
			f.registerProjector(plugin.(comps.IProjector))
		}
	}
}

func (f *CmdSpoke) registerProjector(projector comps.IProjector) {
	f.projector = projector
}

func NewCmdFeature(name schema.Name) *CmdSpoke {
	f := &CmdSpoke{
		handlers:   make([]comps.IMediatorReactor, 0),
		responders: make([]comps.IResponder, 0),
		listeners:  make([]comps.IListener, 0),
		projector:  nil,
	}
	base := NewSpoke(name, f.up, f.down, f.registerReactors)
	f.Spoke = base
	return f
}

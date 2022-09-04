package features

import (
	"context"
	"github.com/hashicorp/go-multierror"
)

type (
	CmdFeatureFtor    func() ICmdFeature
	CmdFeatureBuilder func() ICmdFeature
)

type CmdFeature struct {
	*Feature
	projector  IEventProjector
	responders []IHopeResponder
	listeners  []IFactListener
	handlers   []IMediatorSubscriber
}

func (f *CmdFeature) up(ctx context.Context) error {
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
	if len(errors.Errors) > 0 {
		err := multierror.Flatten(&errors)
		return err
	}
	return nil
}

func (f *CmdFeature) down(ctx context.Context) {
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
	if len(errors.Errors) > 0 {
		err := multierror.Flatten(&errors)
		f.GetLogger().Error(err)
	}
}

func (f *CmdFeature) registerPlugins(plugins []IFeaturePlugin) {
	if len(plugins) == 0 {
		return
	}
	for _, plugin := range plugins {
		switch plugin.(type) {
		case IMediatorSubscriber, IEvtHandler:
			f.handlers = append(f.handlers, plugin.(IMediatorSubscriber))
		case IHopeResponder:
			f.responders = append(f.responders, plugin.(IHopeResponder))
		case IFactListener:
			f.listeners = append(f.listeners, plugin.(IFactListener))
		case IEventProjector:
			f.registerProjector(plugin.(IEventProjector))
		}
	}
}

func (f *CmdFeature) registerProjector(projector IEventProjector) {
	f.projector = projector
}

func NewCmdFeature(name Name) *CmdFeature {
	f := &CmdFeature{
		handlers:   make([]IMediatorSubscriber, 0),
		responders: make([]IHopeResponder, 0),
		listeners:  make([]IFactListener, 0),
		projector:  nil,
	}
	base := NewFeature(name, f.up, f.down, f.registerPlugins)
	f.Feature = base
	return f
}

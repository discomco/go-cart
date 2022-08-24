package features

import (
	"context"

	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/dtos"
	"github.com/hashicorp/go-multierror"
)

type (
	GenCmdFeatureBuilder[T domain.ICmd] func() IGenCmdFeature[T]
	GenCmdFeatureFtor[T domain.ICmd]    func() IGenCmdFeature[T]
)

type (
	GenResponderFtor[THope dtos.IHope] func() IGenResponder[THope]
	IGenResponder[THope dtos.IHope]    interface {
		IHopeResponder
	}
)

type GenCmdFeature struct {
	*Feature
	responders []IHopeResponder
	listeners  []IFactListener
	handlers   []IMediatorSubscriber
}

func (f *GenCmdFeature) up(ctx context.Context) error {
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

func (f *GenCmdFeature) down(ctx context.Context) {
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

func (f *GenCmdFeature) registerCmdPlugins(plugins []IFeaturePlugin) {
	if len(plugins) == 0 {
		return
	}
	for _, plugin := range plugins {
		switch plugin.(type) {
		case IHopeResponder:
			f.responders = append(f.responders, plugin.(IHopeResponder))
		case IFactListener:
			f.listeners = append(f.listeners, plugin.(IFactListener))
		case IMediatorSubscriber:
			f.handlers = append(f.handlers, plugin.(IMediatorSubscriber))
		}
	}
}

func NewGenCmdFeature(name Name) *GenCmdFeature {
	f := &GenCmdFeature{
		handlers:   make([]IMediatorSubscriber, 0),
		responders: make([]IHopeResponder, 0),
		listeners:  make([]IFactListener, 0),
	}
	base := NewFeature(name, f.up, f.down, f.registerCmdPlugins)
	f.Feature = base
	return f
}

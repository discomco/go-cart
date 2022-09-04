package features

import (
	"context"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/reactors"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/hashicorp/go-multierror"
)

type (
	GenCmdSpokeBuilder[T behavior.ICmd] func() IGenCmdSpoke[T]
	GenCmdSpokeFtor[T behavior.ICmd]    func() IGenCmdSpoke[T]
)

type GenCmdSpoke struct {
	*Spoke
	responders []reactors.IResponder
	listeners  []reactors.IListener
	handlers   []reactors.IMediatorReactor
}

func (f *GenCmdSpoke) up(ctx context.Context) error {
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

func (f *GenCmdSpoke) down(ctx context.Context) {
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

func (f *GenCmdSpoke) registerReactors(plugins []reactors.IReactor) {
	if len(plugins) == 0 {
		return
	}
	for _, plugin := range plugins {
		switch plugin.(type) {
		case reactors.IResponder:
			f.responders = append(f.responders, plugin.(reactors.IResponder))
		case reactors.IListener:
			f.listeners = append(f.listeners, plugin.(reactors.IListener))
		case reactors.IMediatorReactor:
			f.handlers = append(f.handlers, plugin.(reactors.IMediatorReactor))
		}
	}
}

func NewGenCmdSpoke(name schema.Name) *GenCmdSpoke {
	f := &GenCmdSpoke{
		handlers:   make([]reactors.IMediatorReactor, 0),
		responders: make([]reactors.IResponder, 0),
		listeners:  make([]reactors.IListener, 0),
	}
	base := NewSpoke(name, f.up, f.down, f.registerReactors)
	f.Spoke = base
	return f
}

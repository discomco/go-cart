package tui

import (
	"github.com/discomco/go-cart/features"
	"github.com/discomco/go-cart/model"
)

type IController interface {
	Register(topic string, action interface{}, transactional bool)
}

type Controller struct {
	*features.AppComponent
	view  IView
	model model.IReadModel
}

func NewController(name features.Name, view IView, model model.IReadModel) *Controller {
	c := &Controller{
		view:  view,
		model: model,
	}
	b := features.NewAppComponent(name)
	c.AppComponent = b
	return c
}

func (c *Controller) Register(topic string, action interface{}, transactional bool) {
	c.GetMediator().RegisterAsync(topic, action, transactional)
	c.GetMediator().WaitAsync()
}

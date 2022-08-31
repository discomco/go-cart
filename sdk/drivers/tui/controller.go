package tui

import (
	"github.com/discomco/go-cart/sdk/features"
)

type IController interface {
	features.IMediatorSubscriber
	Register(topic string, action interface{}, transactional bool)
	IAmController()
	GetProxy() IProxy
}

type Controller struct {
	*features.MsgHandler
	proxy IProxy
}

func (c *Controller) IAmController() {}

func NewController(name features.Name,
	proxy IProxy,
	msgType features.MsgType,
	onMsg features.OnMsgFunc) *Controller {
	c := &Controller{
		proxy: proxy,
	}
	b := features.NewMsgHandler(msgType, onMsg)
	b.Name = name
	c.MsgHandler = b
	return c
}

func (c *Controller) GetProxy() IProxy {
	return c.proxy
}

func (c *Controller) Register(topic string, action interface{}, transactional bool) {
	c.GetMediator().RegisterAsync(topic, action, transactional)
	c.GetMediator().WaitAsync()
}

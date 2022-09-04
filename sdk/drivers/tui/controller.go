package tui

import (
	"github.com/discomco/go-cart/sdk/reactors"
	"github.com/discomco/go-cart/sdk/schema"
)

type IController interface {
	reactors.IMediatorReactor
	Register(topic string, action interface{}, transactional bool)
	IAmController()
	GetProxy() IProxy
}

type Controller struct {
	*reactors.MsgReactor
	proxy IProxy
}

func (c *Controller) IAmController() {}

func NewController(name schema.Name,
	proxy IProxy,
	msgType schema.MsgType,
	onMsg reactors.OnMsgFunc) *Controller {
	c := &Controller{
		proxy: proxy,
	}
	b := reactors.NewMsgReactor(msgType, onMsg)
	b.Name = name
	c.MsgReactor = b
	return c
}

func (c *Controller) GetProxy() IProxy {
	return c.proxy
}

func (c *Controller) Register(topic string, action interface{}, transactional bool) {
	c.GetMediator().RegisterAsync(topic, action, transactional)
	c.GetMediator().WaitAsync()
}

package tui

import (
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
)

type IController interface {
	comps.IMediatorReaction
	Register(topic string, action interface{}, transactional bool)
	IAmController()
	GetProxy() IProxy
}

type Controller struct {
	*comps.MsgReaction
	proxy IProxy
}

func (c *Controller) IAmController() {}

func NewController(name schema.Name,
	proxy IProxy,
	msgType schema.MsgType,
	onMsg comps.OnMsgFunc) *Controller {
	c := &Controller{
		proxy: proxy,
	}
	b := comps.NewMsgReaction(msgType, onMsg)
	b.Name = name
	c.MsgReaction = b
	return c
}

func (c *Controller) GetProxy() IProxy {
	return c.proxy
}

func (c *Controller) Register(topic string, action interface{}, transactional bool) {
	c.GetMediator().RegisterAsync(topic, action, transactional)
	c.GetMediator().WaitAsync()
}

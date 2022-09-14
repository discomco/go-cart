package pulsar

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
)

type OnPulseFunc func(ctx context.Context) error

type pulsarLink struct {
	*comps.Component
	behaviorID schema.IIdentity
	newCH      comps.CmdHandlerFtor
	topic      string
}

func newPulsarLink(name string, topic string, behID schema.IIdentity, newCH comps.CmdHandlerFtor) *pulsarLink {
	pl := &pulsarLink{
		behaviorID: behID,
		newCH:      newCH,
		topic:      topic,
	}
	b := comps.NewComponent(schema.Name(name))
	pl.Component = b
	return pl
}

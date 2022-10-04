package comps

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type Policy struct {
	*EventReaction
	NewCH CmdHandlerFtor
}

func NewPolicy(
	name schema.Name,
	eventType behavior.EventType,
	onEvt OnEvtFunc,
	newCH CmdHandlerFtor,
) *Policy {
	dl := &Policy{
		NewCH: newCH,
	}
	eh := NewEventReaction(eventType, onEvt)
	dl.EventReaction = eh
	dl.Name = name
	return dl
}

func (h *Policy) IAmPolicy() {}

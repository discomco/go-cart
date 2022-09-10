package comps

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type BehaviorLink struct {
	*EventReaction
	NewCH CmdHandlerFtor
}

func NewBehaviorLink(
	name schema.Name,
	eventType behavior.EventType,
	onEvt OnEvtFunc,
	newCH CmdHandlerFtor,
) *BehaviorLink {
	dl := &BehaviorLink{
		NewCH: newCH,
	}
	eh := NewEventReaction(eventType, onEvt)
	dl.EventReaction = eh
	dl.Name = name
	return dl
}

func (h *BehaviorLink) IAmBehaviorLink() {}

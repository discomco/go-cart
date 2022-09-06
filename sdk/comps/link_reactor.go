package comps

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type LinkReactor struct {
	*EventReactor
	NewCH CmdHandlerFtor
}

func NewLinkReactor(
	name schema.Name,
	eventType behavior.EventType,
	onEvt OnEvtFunc,
	newCH CmdHandlerFtor,
) *LinkReactor {
	dl := &LinkReactor{
		NewCH: newCH,
	}
	eh := NewEventReactor(eventType, onEvt)
	dl.EventReactor = eh
	dl.Name = name
	return dl
}

func (h *LinkReactor) IAmLinkReactor() {}

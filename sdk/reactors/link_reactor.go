package reactors

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/schema"
)

type LinkReactor struct {
	*EventReactor
}

func NewLinkReactor(
	name schema.Name,
	eventType behavior.EventType,
	onEvt OnEvtFunc) *LinkReactor {
	dl := &LinkReactor{}
	eh := NewEventReactor(eventType, onEvt)
	dl.EventReactor = eh
	dl.Name = name
	return dl
}

func (h *LinkReactor) IAmLinkReactor() {}

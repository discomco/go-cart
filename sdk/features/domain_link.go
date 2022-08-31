package features

import (
	"github.com/discomco/go-cart/sdk/domain"
)

type DomainLink struct {
	*EventHandler
}

func NewDomainLink(
	name Name,
	eventType domain.EventType,
	onEvt OnEvtFunc) *DomainLink {
	dl := &DomainLink{}
	eh := NewEventHandler(eventType, onEvt)
	dl.EventHandler = eh
	dl.Name = name
	return dl
}

func (h *DomainLink) IAmDomainLink() {}

package list

import (
	"github.com/discomco/go-cart/sdk/core"
)

const (
	LIST_PREFIX = "gamelst"
	ListId      = "42424242-4242-4242-4242-42424242"
)

func DefaultID() *core.Identity {
	ID, _ := core.NewIdentityFrom(LIST_PREFIX, ListId)
	return ID
}

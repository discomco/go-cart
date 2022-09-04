package list

import "github.com/discomco/go-cart/sdk/schema"

const (
	LIST_PREFIX = "gamelst"
	ListId      = "42424242-4242-4242-4242-42424242"
)

func DefaultID() *schema.Identity {
	ID, _ := schema.NewIdentityFrom(LIST_PREFIX, ListId)
	return ID
}

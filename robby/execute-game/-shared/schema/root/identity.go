package root

import "github.com/discomco/go-cart/sdk/core"

const (
	ID_PREFIX = "execgame"
)

func NewRootID() (*core.Identity, error) {
	return core.NewIdentity(ID_PREFIX)
}

func NewRootIDFromString(id string) (*core.Identity, error) {
	return core.NewIdentityFrom(ID_PREFIX, id)
}
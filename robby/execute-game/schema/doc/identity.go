package doc

import "github.com/discomco/go-cart/sdk/core"

const (
	ID_PREFIX = "execgame"
)

func NewGameID() (*core.Identity, error) {
	return core.NewIdentity(ID_PREFIX)
}

func NewGameIDFromString(id string) (*core.Identity, error) {
	return core.NewIdentityFrom(ID_PREFIX, id)
}

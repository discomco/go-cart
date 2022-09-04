package doc

import "github.com/discomco/go-cart/sdk/schema"

const (
	ID_PREFIX = "execgame"
)

func NewGameID() (*schema.Identity, error) {
	return schema.NewIdentity(ID_PREFIX)
}

func NewGameIDFromString(id string) (*schema.Identity, error) {
	return schema.NewIdentityFrom(ID_PREFIX, id)
}

package avatar

import "github.com/discomco/go-cart/sdk/core"

const ID_PREFIX = "avatar"

func NewAvatarID() (*core.Identity, error) {
	return core.NewIdentity(ID_PREFIX)
}

func NewAvatarIDFrom(id string) (*core.Identity, error) {
	return core.NewIdentityFrom(ID_PREFIX, id)
}

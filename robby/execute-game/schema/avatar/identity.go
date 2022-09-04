package avatar

import "github.com/discomco/go-cart/sdk/schema"

const ID_PREFIX = "avatar"

func NewAvatarID() (*schema.Identity, error) {
	return schema.NewIdentity(ID_PREFIX)
}

func NewAvatarIDFrom(id string) (*schema.Identity, error) {
	return schema.NewIdentityFrom(ID_PREFIX, id)
}

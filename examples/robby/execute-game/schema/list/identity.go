package list

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

const (
	LIST_PREFIX = "gamelst"
	ListId      = "42424242-4242-4242-4242-424242424242"
)

func DefaultID() (*schema.Identity, error) {
	ID, err := schema.NewIdentityFrom(LIST_PREFIX, ListId)
	if err != nil {
		return nil, errors.Wrapf(err, "(gameList.DefaultID) failed to create identity")
	}
	return ID, err
}

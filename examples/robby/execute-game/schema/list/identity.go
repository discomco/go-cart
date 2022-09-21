package list

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

const (
	LIST_PREFIX = "calclst"
	ListId      = "42424242-4242-4242-4242-424242424242"
)

// DefaultCalcListID returns the ID for the default calculations list
func DefaultCalcListID() (*schema.Identity, error) {
	ID, err := schema.NewIdentityFrom(LIST_PREFIX, ListId)
	if err != nil {
		return nil, errors.Wrapf(err, "(gameList.DefaultCalcListID) failed to create identity")
	}
	return ID, err
}

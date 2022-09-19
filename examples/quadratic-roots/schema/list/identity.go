package list

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
)

const (
	Prefix = "quadraticlst"
	Id     = "42424242-4242-4242-4242-424242424242"
)

func DefaultID() (*schema.Identity, error) {
	ID, err := schema.NewIdentityFrom(Prefix, Id)
	if err != nil {
		return nil, errors.Wrapf(err, "(quadraticList.DefaultID) failed to create identity")
	}
	return ID, err
}

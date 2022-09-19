package doc

import "github.com/discomco/go-cart/sdk/schema"

const (
	IdPrefix = "quadratic"
)

// NewDocID creates a new Identity for the document, based on the IdPrefix
func NewDocID() (*schema.Identity, error) {
	return schema.NewIdentity(IdPrefix)
}

// NewDocIDFromString takes a string and attempts to create a new identity from it.
func NewDocIDFromString(id string) (*schema.Identity, error) {
	return schema.NewIdentityFrom(IdPrefix, id)
}

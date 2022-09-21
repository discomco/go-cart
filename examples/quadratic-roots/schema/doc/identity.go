package doc

import "github.com/discomco/go-cart/sdk/schema"

const (
	IdPrefix = "quadratic"
)

// NewCalculationID creates a new Identity for the document, based on the IdPrefix
func NewCalculationID() (*schema.Identity, error) {
	return schema.NewIdentity(IdPrefix)
}

// NewCalculationIDFromString takes a string and attempts to create a new identity from it.
func NewCalculationIDFromString(id string) (*schema.Identity, error) {
	return schema.NewIdentityFrom(IdPrefix, id)
}

package contract

import "github.com/discomco/go-cart/sdk/contract"

// IFact is the injector for Generic Listeners
type IFact interface {
	contract.IFact
}

// NewFact returns a new instance of the IFact
func NewFact(calcId string, pl FactPayload) (IFact, error) {
	return contract.NewFact(calcId, pl)
}

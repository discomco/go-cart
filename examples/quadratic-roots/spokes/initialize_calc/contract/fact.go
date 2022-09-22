package contract

import "github.com/discomco/go-cart/sdk/contract"

type IFact interface {
	contract.IFact
}

func NewFact(calcId string, pl Payload) (IFact, error) {
	return contract.NewFact(calcId, pl)
}

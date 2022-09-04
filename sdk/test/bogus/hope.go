package bogus

import "github.com/discomco/go-cart/sdk/contract"

type IHope interface {
	contract.IHope
}

func NewHope(aggregateId string, payload *Car) (IHope, error) {
	return contract.NewHope(aggregateId, payload)
}

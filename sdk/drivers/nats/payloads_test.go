package nats

import (
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/test/bogus"
)

type IHope interface {
	contract.IHope
}

func NewHope(aggregateId string, payload *bogus.Car) (IHope, error) {
	return contract.NewHope(aggregateId, payload)
}

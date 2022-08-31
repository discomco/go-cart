package nats

import (
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/test/bogus"
)

type IHope interface {
	dtos.IHope
}

func NewHope(aggregateId string, payload *bogus.Car) (IHope, error) {
	return dtos.NewHope(aggregateId, payload)
}

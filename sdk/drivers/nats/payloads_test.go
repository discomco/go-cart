package nats

import (
	"github.com/discomco/go-cart/dtos"
	"github.com/discomco/go-cart/test/bogus"
)

type IHope interface {
	dtos.IHope
}

func NewHope(aggregateId string, payload *bogus.Car) (IHope, error) {
	return dtos.NewHope(aggregateId, payload)
}

package bogus

import "github.com/discomco/go-cart/dtos"

type IHope interface {
	dtos.IHope
}

func NewHope(aggregateId string, payload *Car) (IHope, error) {
	return dtos.NewHope(aggregateId, payload)
}

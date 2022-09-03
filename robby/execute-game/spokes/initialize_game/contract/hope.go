package contract

import "github.com/discomco/go-cart/sdk/dtos"

type IHope interface {
	dtos.IHope
}

func NewHope(aggId string, payload Payload) (IHope, error) {
	return dtos.NewHope(aggId, payload)
}

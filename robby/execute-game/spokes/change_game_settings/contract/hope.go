package contract

import "github.com/discomco/go-cart/sdk/contract"

type IHope interface {
	contract.IHope
}

func NewHope(aggId string, payload Payload) (IHope, error) {
	return contract.NewHope(aggId, payload)
}

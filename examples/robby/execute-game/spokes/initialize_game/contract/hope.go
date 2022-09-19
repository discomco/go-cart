package contract

import "github.com/discomco/go-cart/sdk/contract"

type IHope interface {
	contract.IHope
}

func NewHope(behId string, payload Payload) (IHope, error) {
	return contract.NewHope(behId, payload)
}

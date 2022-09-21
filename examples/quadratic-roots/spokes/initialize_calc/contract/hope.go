package contract

import sdk_contract "github.com/discomco/go-cart/sdk/contract"

type IHope interface {
	sdk_contract.IHope
}

func NewHope(calcId string, payload *Payload) (IHope, error) {
	return sdk_contract.NewHope(calcId, payload)
}

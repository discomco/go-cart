package contract

import "github.com/discomco/go-cart/sdk/schema"

type IReq interface {
	IDto
}

type Req struct {
	*Dto
}

func newReq(reqId string, payload schema.IPayload) (*Req, error) {
	d, err := NewDto(reqId, payload)
	if err != nil {
		return nil, err
	}
	q := &Req{}
	q.Dto = d
	return q, nil
}

func NewReq(reqId string, payload schema.IPayload) (*Req, error) {
	return newReq(reqId, payload)
}

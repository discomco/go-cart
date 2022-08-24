package dtos

import "github.com/discomco/go-cart/model"

type IReq interface {
	IDto
}

type Req struct {
	*Dto
}

func newReq(reqId string, payload model.IPayload) (*Req, error) {
	d, err := NewDto(reqId, payload)
	if err != nil {
		return nil, err
	}
	q := &Req{}
	q.Dto = d
	return q, nil
}

func NewReq(reqId string, payload model.IPayload) (*Req, error) {
	return newReq(reqId, payload)
}

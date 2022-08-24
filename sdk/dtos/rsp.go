package dtos

import "github.com/discomco/go-cart/model"

type IRsp interface {
	IFbk
}

type Rsp struct {
	*Fbk
}

func newRsp(reqId string, payload model.IPayload) (*Rsp, error) {
	d := NewFbk(reqId, -1, "")
	r := &Rsp{}
	r.Fbk = d
	err := r.SetJsonData(payload)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func NewRsp(reqId string, payload model.IPayload) (*Rsp, error) {
	return newRsp(reqId, payload)
}

package contract

import "github.com/discomco/go-cart/sdk/schema"

type IRsp interface {
	IFbk
}

type Rsp struct {
	*Fbk
}

func newRsp(reqId string, payload schema.IPayload) (*Rsp, error) {
	d := NewFbk(reqId, -1, "")
	r := &Rsp{}
	r.Fbk = d.(*Fbk)
	err := r.SetPayload(payload)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func NewRsp(reqId string, payload schema.IPayload) (*Rsp, error) {
	return newRsp(reqId, payload)
}

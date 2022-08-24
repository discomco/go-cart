package domain

import (
	"github.com/discomco/go-cart/core"
	"github.com/discomco/go-cart/model"
)

type domainMsg[TID core.IIdentity, TP model.IPayload] struct {
	aID     TID
	payload TP
}

func NewDomainMsg[TID core.IIdentity, TP model.IPayload](aID TID, payload TP) *domainMsg[TID, TP] {
	r := &domainMsg[TID, TP]{aID: aID, payload: payload}
	return r
}
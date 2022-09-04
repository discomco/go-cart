package behavior

import (
	"github.com/discomco/go-cart/sdk/schema"
)

type domainMsg[TID schema.IIdentity, TP schema.IPayload] struct {
	aID     TID
	payload TP
}

func NewDomainMsg[TID schema.IIdentity, TP schema.IPayload](aID TID, payload TP) *domainMsg[TID, TP] {
	r := &domainMsg[TID, TP]{aID: aID, payload: payload}
	return r
}

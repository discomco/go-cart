package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"math/rand"
)

func RandomPayload() *Payload {
	return NewPayload(rand.NormFloat64(), rand.NormFloat64(), rand.NormFloat64())
}

func RandomHope() (IHope, error) {
	ID, _ := doc.NewDocID()
	return NewHope(ID.Id(), RandomPayload())
}

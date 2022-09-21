package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"math/rand"
)

func RandomPayload() *Payload {
	return NewPayload(1_000*rand.NormFloat64(), 1_000*rand.NormFloat64(), 1_000*rand.NormFloat64())
}

func RandomHope() (IHope, error) {
	ID, _ := doc.NewCalculationID()
	return NewHope(ID.Id(), RandomPayload())
}

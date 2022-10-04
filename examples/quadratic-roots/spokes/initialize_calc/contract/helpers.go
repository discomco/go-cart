package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"math/rand"
)

func RandomPayload() *Payload {
	a := 0.0
	b := 0.0
	c := 0.0
	a = 10 * rand.NormFloat64()
	b = 10 * rand.NormFloat64()
	c = 10 * rand.NormFloat64()
	return NewPayload(a, b, c)
}

func RandomHope() (IHope, error) {
	ID, _ := doc.NewCalculationID()
	pl := RandomPayload()
	return NewHope(ID.Id(), *pl)
}

func RandomFact() (IFact, error) {
	ID, _ := doc.NewCalculationID()
	pl := RandomPayload()
	return NewFact(ID.Id(), *pl)
}

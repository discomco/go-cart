package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
)

func RandomHopePayload() *HopePayload {
	return NewHopePayload()
}

func RandomFactPayload() *FactPayload {
	output := schema.RandomOutput(10)
	return NewFactPayload(output)
}

func RandomHope() (IHope, error) {
	ID, _ := doc.NewCalculationID()
	pl := RandomHopePayload()
	return NewHope(ID.Id(), *pl)
}

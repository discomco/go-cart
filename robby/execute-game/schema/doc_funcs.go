package schema

import (
	sdk_model "github.com/discomco/go-cart/sdk/model"
)

type DocFtor sdk_model.DocFtor[GameDoc]

func (r *GameDoc) GetStatus() int {
	return int(r.Status)
}

func GameDocFtor() DocFtor {
	return func() *GameDoc {
		return newGameDoc()
	}
}

func NewGameDoc() *GameDoc {
	return newGameDoc()
}

func newGameDoc() *GameDoc {
	return &GameDoc{
		ID:      nil,
		Details: NewDetails("New GameDoc"),
	}
}

func NewDetails(name string) *Details {
	return &Details{
		Name: name,
	}
}

func NewDimensions(x, y, z int) *Dimensions {
	return &Dimensions{
		X: x,
		Y: y,
		Z: z,
	}
}

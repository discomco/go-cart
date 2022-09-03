package model

import (
	sdk_model "github.com/discomco/go-cart/sdk/model"
)

type DocFtor sdk_model.DocFtor[Root]

func (r *Root) GetStatus() int {
	return int(r.Status)
}

func RootFtor() DocFtor {
	return func() *Root {
		return newRoot()
	}
}

func NewRoot() *Root {
	return newRoot()
}

func newRoot() *Root {
	return &Root{
		ID:      nil,
		Details: NewDetails("New Game"),
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

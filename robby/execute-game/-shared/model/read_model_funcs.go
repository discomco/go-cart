package model

import (
	"github.com/discomco/go-cart/sdk/model"
)

func RootFtor() model.DocFtor[Root] {
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
		Details: nil,
	}
}

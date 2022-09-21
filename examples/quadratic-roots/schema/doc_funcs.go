package schema

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/schema"
)

// GetStatus returns the QuadraticDooc's Status as an integer.
func (doc *QuadraticDoc) GetStatus() int {
	return int(doc.Status)
}

// DocFtor is a functor for QuadraticDoc
func DocFtor() schema.DocFtor[QuadraticDoc] {
	return func() *QuadraticDoc {
		return newDoc()
	}
}

func newDoc() *QuadraticDoc {
	qd := &QuadraticDoc{
		Status: doc.Unknown,
	}
	return qd
}

func NewInput(a float64, b float64, c float64) *Input {
	return &Input{
		A: a,
		B: b,
		C: c,
	}
}

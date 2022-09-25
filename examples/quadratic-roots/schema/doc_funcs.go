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

func NewOutput() *Output {
	return &Output{
		D:  0.0,
		X1: NewImaginaryNumber(0, 0),
		X2: NewImaginaryNumber(0, 0),
	}
}

func NewImaginaryNumber(real float64, imaginary float64) *ImaginaryNumber {
	return &ImaginaryNumber{
		Real:      real,
		Imaginary: imaginary,
	}
}

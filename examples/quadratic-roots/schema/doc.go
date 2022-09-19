package schema

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/schema"
)

// QuadraticDoc is the default read model for the Quadratic roots calculator
type QuadraticDoc struct {
	ID     *schema.Identity
	Input  *Input
	Output *Output
	Status doc.Status
}

// Input is the input for the quadratic calculator and contains the a,b,c coefficients
type Input struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
}

// Output is the output for the quadratic calculator and contains the Discriminator D and the roots X1 and X2
type Output struct {
	D  float64          `json:"d"`
	X1 *ImaginaryNumber `json:"x1"`
	X2 *ImaginaryNumber `json:"x2"`
}

// ImaginaryNumber is a structure that represents an Imaginary number
type ImaginaryNumber struct {
	Real      float64 `json:"real"`
	Imaginary float64 `json:"imaginary"`
}

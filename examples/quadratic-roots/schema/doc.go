package schema

import "github.com/discomco/go-cart/sdk/schema"

type QuadraticDoc struct {
	ID     *schema.Identity
	Input  *Input
	Output *Output
}

// Input is the input for the quadratic calculator and contains the a,b,c coefficients
type Input struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

// Output is the output for the quadratic calculator and contains the Discriminator D and the roots X1 and X2
type Output struct {
	D  int     `json:"d"`
	X1 float64 `json:"x1"`
	X2 float64 `json:"x2"`
}

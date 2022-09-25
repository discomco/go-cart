package schema

import "math/rand"

func RandomInput(factor float64) *Input {
	a := factor * rand.NormFloat64()
	b := factor * rand.NormFloat64()
	c := factor * rand.NormFloat64()
	return NewInput(
		a,
		b,
		c,
	)
}

func RandomOutput(factor float64) *Output {
	o := NewOutput()
	o.D = factor * rand.NormFloat64()
	o.X1 = RandomImaginaryNumber(factor)
	o.X2 = RandomImaginaryNumber(factor)
	return o
}

func RandomImaginaryNumber(factor float64) *ImaginaryNumber {
	return NewImaginaryNumber(
		factor*rand.NormFloat64(),
		factor*rand.NormFloat64(),
	)
}

package behavior

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"math"
)

func calcRoots(input *schema.Input) *schema.Output {
	r := schema.NewOutput()
	a := input.A
	b := input.B
	c := input.C
	r.D = b*b - 4*a*c
	if r.D < 0 {
		r.X1.Real = -b / (2 * a)
		r.X1.Imaginary = math.Sqrt(math.Abs(r.D)) / (2 * a)
		r.X2.Real = -b / (2 * a)
		r.X2.Imaginary = -math.Sqrt(math.Abs(r.D)) / (2 * a)
	} else {
		r.X1.Real = (-b + math.Sqrt(r.D)) / (2 * a)
		r.X2.Real = (-b - math.Sqrt(r.D)) / (2 * a)
	}
	return r
}

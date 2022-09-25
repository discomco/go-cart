package schema

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/schema"
)

// QuadraticList contains a list of Quadratic Root Items
type QuadraticList struct {
	ID    *schema.Identity
	Items map[string]*Calculation
}

// Calculation is the list Item for Quadratic calculations.
type Calculation struct {
	Id            string     `json:"id"`
	Status        doc.Status `json:"status"`
	Equation      string     `json:"equation"`
	Discriminator string     `json:"discriminator"`
	X1            string     `json:"x1"`
	X2            string     `json:"x2"`
}

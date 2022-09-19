package schema

import "github.com/discomco/go-cart/sdk/schema"

// QuadraticList contains a list of Quadratic Root Items
type QuadraticList struct {
	ID    *schema.Identity
	Items map[string]*Calculation
}

// Calculation is the list Item for Quadratic calculations.
type Calculation struct {
	Id            string  `json:"id"`
	Equation      string  `json:"equation"`
	Discriminator float64 `json:"discriminator"`
	Result        string  `json:"result"`
}

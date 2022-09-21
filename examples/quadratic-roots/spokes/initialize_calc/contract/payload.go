package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
)

// Payload is the payload for the initialize_calculation behavior
type Payload struct {
	Input *schema.Input `json:"input"`
}

// NewPayload returns a new payload for the initialize_calculation behavior
func NewPayload(a float64, b float64, c float64) *Payload {
	return &Payload{
		Input: schema.NewInput(a, b, c),
	}
}

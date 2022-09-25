package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
)

// FactPayload is the payload for the calc_roots behavior
type FactPayload struct {
	Output *schema.Output `json:"output"`
}

// NewFactPayload returns a new payload for the calc_roots behavior
func NewFactPayload(output *schema.Output) *FactPayload {
	return &FactPayload{
		Output: output,
	}
}

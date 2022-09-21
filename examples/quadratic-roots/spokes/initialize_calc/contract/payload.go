package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
)

// Payload is the payload for the initialize_calculation behavior
type Payload struct {
	DocId string        `json:"doc_id"`
	Input *schema.Input `json:"input"`
}

// NewPayload returns a new payload for the initialize_calculation behavior
func NewPayload(a float64, b float64, c float64) (*Payload, error) {
	ID, err := doc.NewCalculationID()
	if err != nil {
		return nil, err
	}
	return &Payload{
		DocId: ID.Id(),
		Input: schema.NewInput(a, b, c),
	}, err
}

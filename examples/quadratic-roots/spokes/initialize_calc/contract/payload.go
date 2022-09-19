package contract

// Payload is the payload for the initialize_calculation behavior
type Payload struct {
	A float64 `json:"a,omitempty"`
	B float64 `json:"b,omitempty"`
	C float64 `json:"c,omitempty"`
}

// NewPayload returns a new payload for the initialize_calculation behavior
func NewPayload(a float64, b float64, c float64) *Payload {
	return &Payload{
		A: a,
		B: b,
		C: c,
	}
}

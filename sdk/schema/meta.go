package schema

type Meta struct {
	ID      *Identity
	TraceId string
	Status  int
}

func NewMeta(id *Identity, traceId string, status int) *Meta {
	return &Meta{
		ID:      id,
		TraceId: traceId,
		Status:  status,
	}
}

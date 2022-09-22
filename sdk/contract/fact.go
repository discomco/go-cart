package contract

import (
	"github.com/discomco/go-cart/sdk/schema"
)

//IFact is the injector for DI type discrimination based on Facts
type IFact interface {
	IDto
}

func NewFact(behaviorId string, p schema.IPayload) (*Fact, error) {
	return newFact(behaviorId, p)
}

func newFact(behaviorId string, p schema.IPayload) (*Fact, error) {
	f := &Fact{}
	dto, err := NewDto(behaviorId, p)
	if err != nil {
		return nil, err
	}
	f.Dto = dto
	return f, nil
}

type Fact struct {
	*Dto
}

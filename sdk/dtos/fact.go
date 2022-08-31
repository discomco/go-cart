package dtos

import (
	"github.com/discomco/go-cart/sdk/model"
)

//IFact is the injector for DI type discrimination based on Facts
type IFact interface {
	IDto
}

func NewFact(aggregateId string, p model.IPayload) (*Fact, error) {
	return newFact(aggregateId, p)
}

func newFact(aggregateId string, p model.IPayload) (*Fact, error) {
	f := &Fact{}
	dto, err := NewDto(aggregateId, p)
	if err != nil {
		return nil, err
	}
	f.Dto = dto
	return f, nil
}

type Fact struct {
	*Dto
}

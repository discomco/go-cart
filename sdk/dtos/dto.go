package dtos

import (
	"encoding/json"
	"github.com/discomco/go-cart/sdk/core"
	"github.com/discomco/go-cart/sdk/model"
)

type IDto interface {
	SetJsonData(payload model.IPayload) error
	GetId() string
	GetJsonData(payload model.IPayload) error
	GetID() (core.IIdentity, error)
	SetData(data []byte) IDto
	GetData() []byte
}

type Dto struct {
	Id   string `json:"id"`
	Data []byte `json:"data"`
}

// GetData The Data attached to the Event serialized to bytes.
func (d *Dto) GetData() []byte {
	return d.Data
}

// SetData add the Data attached to the Event serialized to bytes.
func (d *Dto) SetData(data []byte) IDto {
	d.Data = data
	return d
}

func (d *Dto) GetID() (core.IIdentity, error) {
	return core.IdentityFromPrefixedId(d.Id)
}

func (d *Dto) GetJsonData(payload model.IPayload) error {
	err := json.Unmarshal(d.Data, payload)
	if err != nil {
		return err
	}
	return nil
}

// GetAggregateId returns the NewDto's Id as a string
func (d *Dto) GetId() string {
	return d.Id
}

//SetJsonData serializes the given payload and assigns it to the Dto.Data field.
func (d *Dto) SetJsonData(payload model.IPayload) error {
	if payload == nil {
		d.Data = make([]byte, 0)
		return nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	d.Data = data
	return nil
}

func newDto(prefixedId string) *Dto {
	return &Dto{
		Id: prefixedId,
	}

}

//NewDto accepts an Id as a string in the format "prefix-string" and returns an IDto Injector
func NewDto(prefixedId string, payload model.IPayload) (*Dto, error) {
	d := newDto(prefixedId)
	err := d.SetJsonData(payload)
	if err != nil {
		return nil, err
	}
	return d, nil
}

package dtos

import (
	"encoding/json"
	"github.com/discomco/go-cart/sdk/model"
	"sync"
)

type Data2HopeFunc func(data []byte) (IHope, error)

type GenData2HopeFunc[THope IHope] func([]byte) (THope, error)

type HopeType string

type IHope interface {
	IDto
}

type Hope struct {
	*Dto
}

func newHope(aggregateId string, payload model.IPayload) (*Hope, error) {
	res := &Hope{}
	dto, err := NewDto(aggregateId, payload)
	if err != nil {
		return nil, err
	}
	res.Dto = dto
	return res, nil
}

func NewHope(aggregateId string, payload model.IPayload) (*Hope, error) {
	return newHope(aggregateId, payload)
}

var cM = &sync.Mutex{}

func Data2Hope[THope IHope]() GenData2HopeFunc[THope] {
	return func(data []byte) (THope, error) {
		cM.Lock()
		defer cM.Unlock()
		var h THope
		err := json.Unmarshal(data, h)
		if err == nil {
			return h, err
		}
		return h, nil
	}
}

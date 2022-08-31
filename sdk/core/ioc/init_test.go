package ioc

import (
	"github.com/discomco/go-cart/sdk/model"
	"github.com/discomco/go-cart/sdk/test/bogus"
)

var (
	testEnv IDig
)

func init() {
	testEnv = buildTestEnv()
}

type ICar interface {
	model.IWriteModel
}

type IHuman interface {
	model.IWriteModel
}

type HumansAndCars struct {
	list []model.IPayload
}

type Human struct {
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Weight float64 `json:"weight"`
}

func newCar() model.IPayload {
	return bogus.NewCar("Toyota", "Desire", 15, 500)
}

func newHuman() model.IPayload {
	return &Human{
		Name:   "John",
		Age:    42,
		Weight: 50,
	}
}

func NewHumansAndCars(list ...model.IPayload) *HumansAndCars {
	return newHC(list)
}

func newHC(list []model.IPayload) *HumansAndCars {
	return &HumansAndCars{
		list: list,
	}
}

func buildTestEnv() IDig {
	dig := SingleIoC()
	dig.Inject(dig,
		newHuman,
		newCar,
		NewHumansAndCars)
	return dig
}

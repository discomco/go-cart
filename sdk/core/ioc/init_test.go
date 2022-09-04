package ioc

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-cart/sdk/test/bogus"
)

var (
	testEnv IDig
)

func init() {
	testEnv = buildTestEnv()
}

type ICar interface {
	schema.IWriteModel
}

type IHuman interface {
	schema.IWriteModel
}

type HumansAndCars struct {
	list []schema.IPayload
}

type Human struct {
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Weight float64 `json:"weight"`
}

func newCar() schema.IPayload {
	return bogus.NewCar("Toyota", "Desire", 15, 500)
}

func newHuman() schema.IPayload {
	return &Human{
		Name:   "John",
		Age:    42,
		Weight: 50,
	}
}

func NewHumansAndCars(list ...schema.IPayload) *HumansAndCars {
	return newHC(list)
}

func newHC(list []schema.IPayload) *HumansAndCars {
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

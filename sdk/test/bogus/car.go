package bogus

type Car struct {
	Name   string
	Age    int
	Weight float64
	Brand  string
}

func NewCar(brand string, name string, age int, weight float64) *Car {
	return &Car{
		Brand:  brand,
		Name:   name,
		Age:    age,
		Weight: weight,
	}
}

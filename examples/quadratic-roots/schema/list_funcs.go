package schema

import (
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/list"
	"github.com/discomco/go-cart/sdk/schema"
	"log"
	"sync"
)

// NewQuadraticList returns a new empty list of quadratic root calculations
func NewQuadraticList() *QuadraticList {
	Id := QuadraticListKey()
	ID, _ := schema.IdentityFromPrefixedId(Id)
	return &QuadraticList{
		ID:    ID,
		Items: make(map[string]*Calculation),
	}
}

// NewCalculation returns a new Calculation List item
func NewCalculation(Id string, equation string, discriminator float64, result string) *Calculation {
	return &Calculation{
		Id:            Id,
		Equation:      equation,
		Discriminator: discriminator,
		Result:        result,
	}
}

// ListFtor is meant to be injected as a functor to create new QuadraticList instances
func ListFtor() schema.DocFtor[QuadraticList] {
	return func() *QuadraticList {
		return NewQuadraticList()
	}
}

func (l *QuadraticList) GetItem(Id string) *Calculation {
	it, ok := l.Items[Id]
	if !ok {
		it = NewCalculation(Id, "unknown", 0.0, "unknown")
		l.AddItem(it)
	}
	return it
}

var aMutex = &sync.Mutex{}

func (l *QuadraticList) AddItem(item *Calculation) {
	aMutex.Lock()
	defer aMutex.Unlock()
	if l.Items == nil {
		l.Items = make(map[string]*Calculation)
	}
	l.Items[item.Id] = item
}

var dMutex = &sync.Mutex{}

func (l *QuadraticList) DeleteItem(key string) {
	dMutex.Lock()
	defer dMutex.Unlock()
	delete(l.Items, key)
}

func QuadraticListKey() string {
	ID, err := list.DefaultID()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return ID.Id()
}

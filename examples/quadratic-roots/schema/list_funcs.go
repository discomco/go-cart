package schema

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/list"
	"github.com/discomco/go-cart/sdk/schema"
	"log"
	"sync"
)

// NewQuadraticList returns a new empty list of quadratic root calculations
func NewQuadraticList() *QuadraticList {
	Id := DefaultCalcListId()
	ID, _ := schema.IdentityFromPrefixedId(Id)
	return &QuadraticList{
		ID:    ID,
		Items: make(map[string]*Calculation),
	}
}

// NewCalculation returns a new Calculation List item
func NewCalculation(Id string, equation string, discriminator string, result string) *Calculation {
	return &Calculation{
		Id:            Id,
		Equation:      equation,
		Discriminator: discriminator,
		X1:            result,
		Status:        doc.Initialized,
	}
}

// ListFtor is meant to be injected as a functor to create new QuadraticList instances
func ListFtor() schema.DocFtor[QuadraticList] {
	return func() *QuadraticList {
		return NewQuadraticList()
	}
}

// GetItem returns a Calculation List Item.
func (l *QuadraticList) GetItem(Id string) *Calculation {
	it, ok := l.Items[Id]
	if !ok {
		it = NewCalculation(Id, "unknown", "unknown", "unknown")
		l.AddItem(it)
	}
	return it
}

var aMutex = &sync.Mutex{}

// AddItem adds a Calculation Item to the List
func (l *QuadraticList) AddItem(item *Calculation) {
	aMutex.Lock()
	defer aMutex.Unlock()
	if l.Items == nil {
		l.Items = make(map[string]*Calculation)
	}
	l.Items[item.Id] = item
}

var dMutex = &sync.Mutex{}

// DeleteItem deletes a Calculation Item from the list.
func (l *QuadraticList) DeleteItem(key string) {
	dMutex.Lock()
	defer dMutex.Unlock()
	delete(l.Items, key)
}

// DefaultCalcListId returns the List's Identity as a string.
func DefaultCalcListId() string {
	ID, err := list.DefaultCalcListID()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return ID.Id()
}

package bogus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBogusPayload(t *testing.T) {
	// Given
	brand := "Ford"
	name := "Desire"
	age := 25
	weight := 78.5
	// ApplyEvent
	mp := NewCar(brand, name, age, weight)
	// Then
	assert.NotNil(t, mp)
	assert.Equal(t, name, mp.Name)
	assert.Equal(t, age, mp.Age)
	assert.Equal(t, weight, mp.Weight)
}

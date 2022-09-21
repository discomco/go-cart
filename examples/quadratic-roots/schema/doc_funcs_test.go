package schema

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestThatWeCanCreateAQuadraticDoc(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var newDoc schema.DocFtor[QuadraticDoc]
	err := testEnv.Invoke(func(new schema.DocFtor[QuadraticDoc]) {
		newDoc = new
	})
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, newDoc)
	// AND WHEN
	doc := newDoc()
	// THEN
	assert.NotNil(t, doc)
}

func TestThatWeCanCreateANewInput(t *testing.T) {
	// GIVEN
	a := 1_000 * rand.NormFloat64()
	b := 1_000 * rand.NormFloat64()
	c := 1_000 * rand.NormFloat64()
	// WHEN
	ni := NewInput(a, b, c)
	// THEN
	assert.NotNil(t, ni)
	assert.Equal(t, a, ni.A)
	assert.Equal(t, b, ni.B)
	assert.Equal(t, c, ni.C)
}

package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnInitializeCalculationPayload(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	ID, err := doc.NewDocIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	a := 42.0
	b := 4.2
	c := 42.42
	pl := NewPayload(a, b, c)
	// THEN
	assert.NotNil(t, pl)
	assert.Equal(t, a, pl.A)
	assert.Equal(t, b, pl.B)
	assert.Equal(t, c, pl.C)
}

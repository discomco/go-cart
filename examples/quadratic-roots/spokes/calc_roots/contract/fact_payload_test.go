package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnFactPayload(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	ID, err := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	output := schema.RandomOutput(10)
	pl := NewFactPayload(output)
	// THEN
	assert.NotNil(t, pl)
	assert.NotNil(t, pl.Output)
	assert.Equal(t, output.D, pl.Output.D)
	assert.Equal(t, output.X1, pl.Output.X1)
	assert.Equal(t, output.X2, pl.Output.X2)
}

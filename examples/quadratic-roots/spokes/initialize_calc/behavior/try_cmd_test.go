package behavior

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestThatWeCanInitializeACalculation(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestCalculation)
	a := 1_000 * rand.NormFloat64()
	b := 1_000 * rand.NormFloat64()
	c := 1_000 * rand.NormFloat64()
	// WHEN
	calculation := newTestCalculation()
	assert.NotNil(t, calculation)

	ID, err := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	calculation.SetID(ID)

	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()

	pl := contract.NewPayload(a, b, c)

	initCmd, err := NewCmd(ID, *pl)

	evt, fbk := calculation.TryCommand(ctx, initCmd)
	state := calculation.GetState().(*schema.QuadraticDoc)

	assert.NotNil(t, evt)
	assert.NotNil(t, fbk)
	assert.NotNil(t, state.Input)
	assert.Equal(t, a, state.Input.A)
	assert.Equal(t, b, state.Input.B)
	assert.Equal(t, c, state.Input.C)
	assert.True(t, fbk.IsSuccess())
	assert.Equal(t, fbk.GetStatus(), int(state.Status))
}

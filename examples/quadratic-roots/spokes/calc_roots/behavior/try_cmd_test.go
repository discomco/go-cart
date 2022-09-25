package behavior

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/contract"
	initialize_calc_behavior "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	initialize_calc_contract "github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanCalculateRoots(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestCalculation)
	// WHEN
	calculation := newTestCalculation()
	assert.NotNil(t, calculation)

	ID, err := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	calculation.SetID(ID)

	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()

	initPl := initialize_calc_contract.RandomPayload()
	initCmd, _ := initialize_calc_behavior.NewCmd(ID, *initPl)
	initEvt, initFbk := calculation.TryCommand(ctx, initCmd)
	assert.NotNil(t, initEvt)
	assert.True(t, initFbk.IsSuccess())

	pl := contract.NewHopePayload()
	calcCmd, err := NewCmd(ID, *pl)

	calcEvt, calcFbk := calculation.TryCommand(ctx, calcCmd)
	state := calculation.GetState().(*schema.QuadraticDoc)

	assert.NotNil(t, calcEvt)
	assert.NotNil(t, calcFbk)
	assert.NotNil(t, state.Input)
	assert.True(t, calcFbk.IsSuccess())
	assert.Equal(t, calcFbk.GetStatus(), int(state.Status))
}

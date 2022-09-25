package comps

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/calc_roots/behavior"
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
	assert.NotNil(t, testEnv)
	assert.NotNil(t, newTestCH)
	// AND
	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()
	// AND
	initPl := initialize_calc_contract.RandomPayload()
	calcID, _ := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	//	calcID, _ := doc.NewCalculationID()

	initCmd, _ := initialize_calc_behavior.NewCmd(calcID, *initPl)
	// AND
	ch := newTestCH()
	// WHEN
	fbk := ch.Handle(ctx, initCmd)
	if !fbk.IsSuccess() {
		testLogger.Errorf(fbk.GetFlattenedErrors())
		t.Fatal(fbk.GetFlattenedErrors())
	}

	calcRootsPl := contract.RandomHopePayload()
	calcCmd, _ := behavior.NewCmd(calcID, *calcRootsPl)

	ch = newTestCH()
	calcFbk := ch.Handle(ctx, calcCmd)
	if !calcFbk.IsSuccess() {
		testLogger.Errorf(calcFbk.GetFlattenedErrors())
		t.Fatal(calcFbk.GetFlattenedErrors())
	}
	// THEN
	assert.NotNil(t, calcFbk)
	assert.True(t, calcFbk.IsSuccess())
}

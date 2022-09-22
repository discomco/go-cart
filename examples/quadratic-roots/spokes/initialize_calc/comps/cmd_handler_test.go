package comps

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/behavior"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanInitializeACalculation(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, newTestCH)
	// AND
	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()
	// AND
	pl := contract.RandomPayload()
	calcID, _ := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	//	calcID, _ := doc.NewCalculationID()

	initCmd, _ := behavior.NewCmd(calcID, *pl)
	// AND
	ch := newTestCH()
	// WHEN
	fbk := ch.Handle(ctx, initCmd)
	if !fbk.IsSuccess() {
		testLogger.Errorf(fbk.GetFlattenedErrors())
		t.Fatal(fbk.GetFlattenedErrors())
	}
	// THEN
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
}

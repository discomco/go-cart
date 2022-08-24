package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_That_We_Can_Register_An_Actor(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// AND
	var agg IAggregate
	testEnv.Invoke(func(ftor AggFtor) {
		agg = ftor()
	})
	// WHEN
	agg.Inject(agg,
		AnExec,
		AnApply,
	)
	// THEN
	assert.NotNil(t, agg)
}

func Test_That_We_Can_Check_Aggregate_Capabilities(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// AND
	var agg IAggregate
	testEnv.Invoke(func(ftor AggFtor) {
		agg = ftor()
	})
	// WHEN
	agg.Inject(agg,
		AnExec,
		AnApply,
	)
	// WHEN
	knowsCmd := agg.KnowsCmd(A_CMD_TOPIC)
	// THEN
	assert.True(t, knowsCmd)
	// AND WHEN
	knowsEvt := agg.KnowsEvt(A_EVT_TOPIC)
	// THEN
	assert.True(t, knowsEvt)

}

package behavior

import (
	"context"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanInitializeAnAggregate(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestBehavior)
	// WHEN
	agg := newTestBehavior()
	assert.NotNil(t, agg)

	ID, err := doc.NewGameIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	agg.SetID(ID)

	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()
	pl := contract.NewPayload(ID.Id(), "New GameDoc", 42, 42, 42, 12)
	initCmd, err := NewCmd(ID, *pl)

	evt, fbk := agg.TryCommand(ctx, initCmd)
	state := agg.GetState().(*schema.GameDoc)

	assert.NotNil(t, evt)
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
	assert.Equal(t, fbk.GetAggregateStatus(), int(state.Status))
}

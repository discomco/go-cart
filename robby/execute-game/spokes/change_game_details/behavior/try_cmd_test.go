package behavior

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	change_game_details_testing "github.com/discomco/go-cart/robby/execute-game/spokes/change_game_details/testing"
	initialize_game_behavior "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	initialize_game_contract "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanInitializeABehaviorAndChangeItsDetails(t *testing.T) {
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
	// INITIALIZE FIRST
	initPl := initialize_game_contract.NewPayload(ID.Id(), "New GameDoc", 42, 42, 42, 12)
	initCmd, err := initialize_game_behavior.NewCmd(ID, *initPl)
	evt, fbk := agg.TryCommand(ctx, initCmd)
	state := agg.GetState().(*schema.GameDoc)

	assert.NotNil(t, evt)
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
	assert.Equal(t, fbk.GetAggregateStatus(), int(state.Status))

	// CHANGE DETAILS
	changeDetailsPl := change_game_details_testing.RandomPayload()
	assert.NotNil(t, changeDetailsPl)
	changeDetailsCmd, err := NewCmd(ID, *changeDetailsPl)
	assert.NotNil(t, changeDetailsCmd)
	assert.NoError(t, err)
	changeDetailsEvt, changeDetailsFbk := agg.TryCommand(ctx, changeDetailsCmd)
	assert.NotNil(t, changeDetailsEvt)
	assert.NotNil(t, changeDetailsFbk)
	assert.True(t, changeDetailsFbk.IsSuccess())
}

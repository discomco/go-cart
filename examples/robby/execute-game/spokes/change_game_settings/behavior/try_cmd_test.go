package behavior

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/schema"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/change_game_settings/contract"
	initialize_game_behavior "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	initialize_game_contract "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanInitializeABehaviorAndChangeItsSettings(t *testing.T) {
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
	changeSettingsPl := contract.RandomPayload()
	assert.NotNil(t, changeSettingsPl)
	changeSettingsCmd, err := NewCmd(ID, *changeSettingsPl)
	assert.NotNil(t, changeSettingsCmd)
	assert.NoError(t, err)
	changeSettingsEvt, changeSettingsFbk := agg.TryCommand(ctx, changeSettingsCmd)
	assert.NotNil(t, changeSettingsEvt)
	assert.NotNil(t, changeSettingsFbk)
	assert.True(t, changeSettingsFbk.IsSuccess())
}

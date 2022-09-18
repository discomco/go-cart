package comps

import (
	"context"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema"
	"github.com/discomco/go-cart/examples/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/behavior"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	initialize_game "github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/behavior"
	contract2 "github.com/discomco/go-cart/examples/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanResolveACommandHandler(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var newCh comps.CmdHandlerFtor
	err := testEnv.Invoke(func(nc comps.CmdHandlerFtor) {
		newCh = nc
	})
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, newCh)
}

func TestThatWeCanExecuteAChangeEventDetailsCmd(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestCmdHandler)
	// AND
	//	Id, err := root.NewGameIDFromString(test.CLEAN_TEST_UUID)
	ID, err := doc.NewGameID()
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	// AND
	initPl := contract2.NewPayload(ID.Id(), "John's Game", 23, 22, 42, 5)
	initCmd, err := initialize_game.NewCmd(ID, *initPl)
	assert.NoError(t, err)
	assert.NotNil(t, initCmd)

	// AND
	details := schema.NewDetails("Paul's Robby Game")
	// AND
	changeGameDetailsPl := contract.NewPayload(details)
	assert.NotNil(t, changeGameDetailsPl)
	// AND
	changeGameDetailsCmd, err := behavior.NewCmd(ID, *changeGameDetailsPl)
	assert.NoError(t, err)
	assert.NotNil(t, changeGameDetailsCmd)
	// AND
	ch := newTestCmdHandler()
	assert.NotNil(t, ch)
	// AND
	ctx, elapsed := context.WithTimeout(context.Background(), 10*time.Minute)
	defer elapsed()
	assert.NotNil(t, ctx)

	// WHEN
	fbk := ch.Handle(ctx, initCmd)
	// THEN
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
	assert.Equal(t, int(doc.Initialized), fbk.GetAggregateStatus())

}

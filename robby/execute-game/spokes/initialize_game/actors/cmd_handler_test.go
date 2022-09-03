package actors

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/behavior"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanResolveACommandHandler(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var newCh features.CmdHandlerFtor
	err := testEnv.Invoke(func(nc features.CmdHandlerFtor) {
		newCh = nc
	})
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, newCh)
}

func TestThatWeCanExecuteAnInitializeCmd(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestCmdHandler)
	// AND
	//	ID, err := root.NewGameIDFromString(test.CLEAN_TEST_UUID)
	ID, err := doc.NewGameID()
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	// AND
	pl := contract.NewPayload(ID.Id(), "John's Robby Game", 42, 42, 42, 12)
	assert.NotNil(t, pl)
	// AND
	initCmd, err := behavior.NewCmd(ID, *pl)
	assert.NoError(t, err)
	assert.NotNil(t, initCmd)
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

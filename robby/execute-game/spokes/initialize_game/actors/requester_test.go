package actors

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	testing2 "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/testing"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanResolveARequester(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	err := testEnv.Invoke(func(requester IRequester) {

	})
	assert.NoError(t, err)
}

func TestThatWeCanRequestAGameInitialization(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testRequester)
	// AND
	ID, err := doc.NewGameID()
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	// AND
	hope, err := testing2.RandomHope()
	assert.NoError(t, err)
	assert.NotNil(t, hope)
	// AND
	ctx, elapsed := context.WithTimeout(context.Background(), 10*time.Second)
	defer elapsed()
	assert.NotNil(t, ctx)
	// WHEN
	fbk := testRequester.Request(ctx, hope, 10*time.Second)
	// THEN
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
}

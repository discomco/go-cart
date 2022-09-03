package domain

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"
	"github.com/discomco/go-cart/robby/execute-game/-shared/model/root"
	"github.com/discomco/go-cart/robby/execute-game/modules/initialize/dtos"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanInitializeAnAggregate(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestAgg)
	// WHEN
	agg := newTestAgg()
	assert.NotNil(t, agg)

	ID, err := root.NewRootIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	agg.SetID(ID)

	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()
	pl := dtos.NewPayload("New Game", 42, 42, 42, 12)
	initCmd, err := NewCmd(ID, *pl)

	evt, fbk := agg.TryCommand(ctx, initCmd)
	state := agg.GetState().(*model.Root)

	assert.NotNil(t, evt)
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
	assert.Equal(t, fbk.GetAggregateStatus(), int(state.Status))
}

package contract

import (
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	testing2 "github.com/discomco/go-cart/robby/execute-game/spokes/change_game_settings/testing"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnSettingsPayload(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	ID, err := doc.NewGameIDFromString(test.CLEAN_TEST_UUID)
	assert.NoError(t, err)
	assert.NotNil(t, ID)
	pl := testing2.RandomPayload()
	// THEN
	assert.NotNil(t, pl)
	assert.NotNil(t, pl.Settings)
}

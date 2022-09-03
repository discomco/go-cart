package actors

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

}


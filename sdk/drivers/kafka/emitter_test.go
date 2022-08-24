package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAnEmitter(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	e, err := newEmitter("test.KafkaEmitter", "", nil)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, e)
}

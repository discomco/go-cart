package reactors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewEvtHandler(t *testing.T) {
	// GIVEN
	ioc := buildTestEnv()
	assert.NotNil(t, ioc)
	// WHEN
	eh := NewEventReactor("-base", nil)
	// THEN
	assert.NotNil(t, eh)
}

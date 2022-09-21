package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanGetTheDefaultListID(t *testing.T) {
	// GIVEN
	// WHEN
	ID, err := DefaultCalcListID()
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, ID)
}

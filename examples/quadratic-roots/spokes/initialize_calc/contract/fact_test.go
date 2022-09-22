package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateANewFact(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	ID, _ := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	pl := RandomPayload()
	// WHEN
	f, err := NewFact(ID.Id(), *pl)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

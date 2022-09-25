package contract

import (
	"github.com/discomco/go-cart/examples/quadratic-roots/schema/doc"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateANewHope(t *testing.T) {
	// GIVEN
	pl := RandomHopePayload()
	ID, _ := doc.NewCalculationIDFromString(test.CLEAN_TEST_UUID)
	// WHEN
	h, _ := NewHope(ID.Id(), *pl)
	// THEN
	assert.NotNil(t, h)
}

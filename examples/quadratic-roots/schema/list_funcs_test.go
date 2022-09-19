package schema

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAList(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testLogger)
	// WHEN
	var list *QuadraticList
	err := testEnv.Invoke(func(newList schema.DocFtor[QuadraticList]) {
		list = newList()
	})
	if err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	// THEN
	assert.NotNil(t, list)
}

package schema

import (
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanCreateAGameDoc(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var newDoc schema.DocFtor[GameDoc]
	err := testEnv.Invoke(func(new schema.DocFtor[GameDoc]) {
		newDoc = new
	})
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, newDoc)
	// AND WHEN
	doc := newDoc()
	// THEN
	assert.NotNil(t, doc)
}

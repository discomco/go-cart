package features

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func TestThatWeCanCreateAGenericEventHandler(t *testing.T) {
	// GIVEN
	// WHEN
	h := newMyGenHandler()
	// THEN
	assert.NotNil(t, h)
}

func TestThatWeCanHandleAnEvent(t *testing.T) {
	// GIVEN
	_ = buildTestEnv()
	// WHEN
	h := newMyGenHandler()
	assert.NotNil(t, h)
	// AND WE CREATE AN EVENT
	evt, err := newMyTestEvt()
	assert.NoError(t, err)
	assert.NotNil(t, evt)
	// AND WE SUPPLY A CONTEXT
	ctx := context.Background()
	assert.NotNil(t, ctx)
	// WHEN
	err = h.When(ctx, evt)
	// THEN
	assert.Nil(t, err)
}

func TestThatWeCanInjectAGenericEventHandlerConstructorAndRetrieveIt(t *testing.T) {
	// Given
	ioc := buildTestEnv()
	assert.NotNil(t, ioc)

}

package comps

import (
	"context"
	"github.com/discomco/go-cart/examples/quadratic-roots/spokes/initialize_calc/contract"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThatWeCanCreateARequester(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var requester IRequester
	if err := testEnv.Invoke(func(ftor comps.GenRequesterFtor[contract.IHope]) {
		requester, _ = ftor()
	}); err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	// THEN
	assert.NotNil(t, requester)
}

func TestThatWeCanSendARequestUsingTheRequester(t *testing.T) {
	// GIVEN
	assert.NotNil(t, newTestRequester)
	// AND
	requester, err := newTestRequester()
	if err != nil {
		testLogger.Fatal(err)
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, requester)
	// AND
	ctx, expired := context.WithTimeout(context.Background(), 10*time.Second)
	defer expired()
	assert.NotNil(t, ctx)
	// AND
	hope, err := contract.RandomHope()
	assert.NoError(t, err)
	assert.NotNil(t, hope)
	// WHEN
	fbk := requester.RequestAsync(ctx, hope, 10*time.Second)
	// THEN
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
	if !fbk.IsSuccess() {
		testLogger.Error(fbk.GetFlattenedErrors())
	}
}

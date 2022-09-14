package pulsar

import (
	"context"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestThatWeCanCreateAPulsar(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	name := "test.pulsar"
	topic := "test:pulsar"
	interval := 10 * time.Second
	// WHEN
	p := newPulsar(schema.Name(name), topic, interval)
	// THEN
	assert.NotNil(t, p)
	assert.Equal(t, name, string(p.Name))
}

func TestThatWeCanResolveASecondPulsar(t *testing.T) {
	//  GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var pulsar IPulsar
	err := testEnv.Invoke(func(p ISecondPulsar) {
		pulsar = p
	})
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, pulsar)
}

func TestThatWeCanActivateAndDeactivateASecondPulsar(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	assert.NotNil(t, testSecondPulsar)
	assert.NotNil(t, testMediatorLogger)
	// AND
	ctx, timedOut := context.WithTimeout(context.Background(), 5*time.Second)
	defer timedOut()
	assert.NotNil(t, ctx)
	// AND
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(runPulsar(ctx))
	eg.Go(runMediatorLogger(ctx))
	err := eg.Wait()
	// THEN
	assert.NoError(t, err)
}

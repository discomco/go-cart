package probes

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestThatWeCanInstantiateAMetricsService(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var bm IBogusMetrics
	err := testEnv.Invoke(func(m IBogusMetrics) {
		bm = m
	})
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, bm)
}

func TestThatWeCanRunAMetricsService(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// AND
	ctxR, done := context.WithTimeout(context.Background(), 10*time.Second)
	defer done()
	assert.NotNil(t, ctxR)
	// AND
	var bm IBogusMetrics
	err := testEnv.Invoke(func(m IBogusMetrics) {
		bm = m
	})
	assert.NoError(t, err)
	assert.NotNil(t, bm)
	// WHEN
	gr, ctx := errgroup.WithContext(ctxR)
	gr.Go(bm.Run(ctx))
	err = gr.Wait()
	assert.NoError(t, err)
	bm.Shutdown(ctx)
}

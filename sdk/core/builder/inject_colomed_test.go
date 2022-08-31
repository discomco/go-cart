package builder

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/discomco/go-cart/sdk/core/logger"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatWeCanResolveALogger(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var log logger.IAppLogger
	err := testEnv.Invoke(func(logger logger.IAppLogger) {
		log = logger
	})
	assert.NoError(t, err)
	// THEN
	assert.NotNil(t, log)
}

func TestThatWeCanResolveAnAppConfig(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var cfg config.IAppConfig
	err := testEnv.Invoke(func(ac config.IAppConfig) {
		cfg = ac
	})
	assert.NoError(t, err)
	// THEN
	assert.NotNil(t, cfg)
}

func TestThatWeCanResolveAMediator(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var med mediator.IMediator
	err := testEnv.Invoke(func(m mediator.IMediator) {
		med = m
	})
	assert.NoError(t, err)
	// THEN
	assert.NotNil(t, med)
}

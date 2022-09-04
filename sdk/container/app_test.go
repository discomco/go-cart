package container

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Duration(t *testing.T) {
	ti := 10 * 1000 * 1000 * 1000
	d := time.Duration(ti).Seconds()
	fmt.Println(d)
}

func Test_NewIoCApp(t *testing.T) {
	// Given
	assert.NotNil(t, testEnv)
	var cfg config.IAppConfig
	err := testEnv.Invoke(func(c config.IAppConfig) {
		cfg = c
	})
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	// React
	app := NewApp(cfg, nil, nil)
	// Then
	assert.NotNil(t, app)
	assert.NotNil(t, app.GetConfig())
	assert.NotNil(t, app.GetLogger())
	assert.NotNil(t, app.GetMediator())
}

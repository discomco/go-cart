package spoke

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/actors"
	testing2 "github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/testing"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestThatWeCanResolveAModule(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// AND
	var module ISpoke
	err := testEnv.Invoke(func(mod ISpoke) {
		module = mod
	})
	assert.NoError(t, err)
	assert.NotNil(t, module)
}

func TestThatWeCanResolveARequester(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	err := testEnv.Invoke(func(_ actors.IRequester) {

	})
	assert.NoError(t, err)
}

func TestThatWeCanRunASpoke(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testModule)
	assert.NotNil(t, testRequester)
	// AND
	ctx, expired := context.WithTimeout(context.Background(), 20*time.Second)
	defer expired()
	// AND
	fbks := make(chan dtos.IFbk)
	// WHEN

	go func(c context.Context, fc chan dtos.IFbk) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for {
				fbk := <-fc
				assert.True(t, fbk.IsSuccess())
			}
		}
	}(ctx, fbks)

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(testModule.Run(ctx))
	wg.Go(requestWorker(ctx, fbks))
	wg.Wait()

	// THEN

}

func requestWorker(ctx context.Context, fbks chan dtos.IFbk) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for i := 0; i < 20; i++ {
				hope, err := testing2.RandomHope()
				if err != nil {
					return err
				}
				fbk := testRequester.Request(ctx, hope, 10*time.Second)
				fbks <- fbk
			}
			return nil
		}
	}
}

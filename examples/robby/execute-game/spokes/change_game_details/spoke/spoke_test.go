package spoke

import (
	"context"
	"github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/comps"
	change_game_details_contract "github.com/discomco/go-cart/examples/robby/execute-game/spokes/change_game_details/contract"
	"github.com/discomco/go-cart/sdk/contract"
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
	err := testEnv.Invoke(func(_ comps.IRequester) {

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
	fbks := make(chan contract.IFbk)
	// WHEN

	go func(c context.Context, fc chan contract.IFbk) error {
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

func requestWorker(ctx context.Context, fbks chan contract.IFbk) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			for i := 0; i < 20; i++ {
				hope, err := change_game_details_contract.RandomHope()
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

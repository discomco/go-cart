package spoke

import (
	"context"
	"github.com/discomco/go-cart/robby/execute-game/schema/doc"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/actors"
	"github.com/discomco/go-cart/robby/execute-game/spokes/initialize_game/contract"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"sync"
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
				hope, err := randomHope()
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

var rMutex = &sync.Mutex{}

func randomHope() (contract.IHope, error) {
	rMutex.Lock()
	defer rMutex.Unlock()
	aggID, _ := doc.NewGameID()
	pl := randomPayload()
	return contract.NewHope(aggID.Id(), *pl)
}

var (
	gameNames = []string{
		"John's Bonanza",
		"All quiet on the Southern Front",
		"Resurrection",
		"The Day after Yesterday",
		"Corpses for Sale",
	}
)

func randomPayload() *contract.Payload {
	ID, _ := doc.NewGameID()
	seed := rand.Intn(5)
	name := gameNames[seed]
	x := rand.Intn(42) + 3
	y := rand.Intn(42) + 3
	z := rand.Intn(42) + 3
	nbrOfPlayers := rand.Intn(12) + 2
	return contract.NewPayload(ID.Id(), name, x, y, z, nbrOfPlayers)
}

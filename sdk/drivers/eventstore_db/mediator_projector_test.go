package eventstore_db

import (
	"github.com/discomco/go-cart/core"
	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/features"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestThatWeCanCreateAProjector(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// WHEN
	var prj features.IEventProjector
	err := testEnv.Invoke(func(prjCtor features.EventProjectorFtor) {
		prj = prjCtor()
	})
	// THEN
	assert.Nil(t, err)
	assert.NotNil(t, prj)
}

func TestThatWeCanProject(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	ctx := context.Background()
	assert.NotNil(t, ctx)
	// AND
	for {
		select {
		case <-ctx.Done():
			break
		default:
			testProjector.Project(ctx,
				[]string{testConfig.GetProjectionConfig().GetEventPrefix()},
				testConfig.GetProjectionConfig().GetPoolSize())
		}
	}
}

func TestThatEventsArriveAtTheMediator(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	ctx, done := context.WithTimeout(context.Background(), 2*time.Minute)
	defer done()

	assert.NotNil(t, ctx)

	assert.NotNil(t, testLoggingHandler)

	err := testLoggingHandler.Activate(ctx)
	assert.NoError(t, err)
	// WHEN
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(pusherWorker(ctx))

	grp.Go(projectorWorker(ctx))

	err = grp.Wait()
	assert.Nil(t, err)
}

func TestThatWhenWorks(t *testing.T) {
	assert.NotNil(t, testProjector)
	ctx, done := context.WithTimeout(context.Background(), 2*time.Minute)
	defer done()
	agg := domain.NewAggregate("test", "1")
	ID, _ := core.NewIdentity("test")
	agg.SetID(ID)
	testEvent := domain.NewEvt(agg, "test")

	testMed.Register(testEvent.GetEventTypeString(), func(ctx context.Context, evt domain.IEvt) {
		assert.Equal(t, testEvent, evt)
	})

	// WHEN
	err := testProjector.When(ctx, testEvent)
	assert.NoError(t, err)

}

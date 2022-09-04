package eventstore_db

import (
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/reactors"
	"github.com/discomco/go-cart/sdk/schema"
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
	var prj reactors.IProjector
	err := testEnv.Invoke(func(prjCtor reactors.ProjectorFtor) {
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
	agg := behavior.NewBehavior("test", nil)
	ID, _ := schema.NewIdentity("test")
	agg.SetID(ID)
	testEvent := behavior.NewEvt(agg, "test")

	testMed.Register(testEvent.GetEventTypeString(), func(ctx context.Context, evt behavior.IEvt) {
		assert.Equal(t, testEvent, evt)
	})

	// WHEN
	err := testProjector.React(ctx, testEvent)
	assert.NoError(t, err)

}

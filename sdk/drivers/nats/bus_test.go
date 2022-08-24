package nats

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestNewBus(t *testing.T) {
	// Given
	ioc := buildTestEnv()
	// When
	assert.NotNil(t, ioc)
}

func TestThatNATSBusImplementsNatsIBus(t *testing.T) {
	// Given
	ioc := buildTestEnv()
	var b INATSBus
	var err error
	err = ioc.Invoke(func(ctor BusFtor) {
		b, err = ctor()
	})
	assert.Nil(t, err)
	assert.NotNil(t, b)
	// When
	ok := IAmNATSBus(b)
	// Then
	assert.True(t, ok)
}

func TestThatWeCanListenOnATopic(t *testing.T) {
	// GIVEN
	assert.NotNil(t, testEnv)
	// AND
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	assert.NotNil(t, ctx)
	// AND
	var b INATSBus
	var err error
	testEnv.Invoke(func(ctor BusFtor) {
		b, err = ctor()
	})
	assert.NoError(t, err)
	assert.NotNil(t, b)
	//AND
	// WHEN
	outs := make(chan []byte)
	ins := make(chan []byte)
	// AND
	g, ctx := errgroup.WithContext(ctx)
	g.Go(listenWorker(ctx, outs))
	g.Go(publishWorker(ctx, ins))
	// AND
	go func() {

		for {
			select {
			case <-ctx.Done():
			default:
				{
					time.Sleep(3 * time.Second)
					testLogger.Infof("Sending message Hello World")
					ins <- []byte("Hello World")
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case data := <-outs:
			{
				testLogger.Infof("Received: %s", string(data))
				//assert.Equal(t, []byte("Hello World"), data)
			}
		}
	}
	err = g.Wait()
	// Then
	assert.NotNil(t, err)
}

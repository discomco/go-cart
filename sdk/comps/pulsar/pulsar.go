package pulsar

import (
	"context"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/sync/errgroup"
	"time"
)

const (
	OnSecond = "pulsar.OnSecond"
	OnMinute = "pulsar.OnMinute"
	OnHour   = "pulsar.OnHour"
)

type IHourPulsar interface {
	IPulsar
}

type IMinutePulsar interface {
	IPulsar
}

type ISecondPulsar interface {
	IPulsar
}

// IPulsar is an active component that triggers a pulse at a certain interval and publishes it onto the Mediator.
type IPulsar interface {
	comps.ISpokePlugin
	IAmPulsar()
}

type pulsar struct {
	*comps.Component
	topic  string
	ticker *time.Ticker
}

func (p *pulsar) tickerWorker(ctx context.Context) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case t := <-p.ticker.C:
				p.GetMediator().Broadcast(p.topic, ctx, t)
			}
		}
	}
}

func (p *pulsar) Activate(ctx context.Context) error {
	p.GetLogger().Infof("Activating Pulsar %+v", p.topic)
	p.GetMediator().RegisterTopic(p.topic)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(p.tickerWorker(ctx))
	return eg.Wait()
}

func (p *pulsar) Deactivate(ctx context.Context) error {
	p.GetLogger().Infof("Deactivating Pulsar: [%+v]", p.topic)
	p.GetMediator().UnregisterTopic(p.topic)
	p.ticker.Stop()
	return nil
}

func (p *pulsar) IAmPulsar() {}

func HourPulsar() IHourPulsar {
	return NewPulsar("pulsar.Hour", OnHour, 1*time.Hour)
}

func MinutePulsar() IMinutePulsar {
	return NewPulsar("pulsar.Minute", OnMinute, 1*time.Minute)
}

func SecondPulsar() ISecondPulsar {
	return NewPulsar("pulsar.Second", OnSecond, 1*time.Second)
}

func NewPulsar(name string, topic string, interval time.Duration) IPulsar {
	return newPulsar(schema.Name(name), topic, interval)
}

func newPulsar(name schema.Name, topic string, interval time.Duration) *pulsar {
	p := &pulsar{
		ticker: time.NewTicker(interval),
		topic:  topic,
	}
	b := comps.NewComponent(name)
	p.Component = b
	return p
}

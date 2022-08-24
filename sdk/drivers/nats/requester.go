package nats

import (
	"context"
	"fmt"
	"time"

	"github.com/discomco/go-cart/core/ioc"
	"github.com/discomco/go-cart/core/utils/convert"
	"github.com/discomco/go-cart/dtos"
	"github.com/discomco/go-cart/features"
	"golang.org/x/sync/errgroup"
)

type INATSRequester[THope dtos.IHope] interface {
	features.IGenHopeRequester[THope]
}

type Requester[THope dtos.IHope] struct {
	*features.AppComponent
	bus   INATSBus
	Topic string
}

func (r *Requester[THope]) IAmHopeRequester() {}

const (
	RequesterFmt = "[%+v].NATSRequester"
)

func newRequester[THope dtos.IHope](topic string) (*Requester[THope], error) {
	name := fmt.Sprintf(RequesterFmt, topic)
	ac := features.NewAppComponent(features.Name(name))
	dig := ioc.SingleIoC()
	var b INATSBus
	var err error
	err = dig.Invoke(func(ftor BusFtor) {
		b, err = ftor()
	})
	if err != nil {
		return nil, err
	}
	r := &Requester[THope]{
		bus:          b,
		AppComponent: ac,
		Topic:        topic,
	}
	return r, nil
}

func (r *Requester[THope]) requestRaw(ctx context.Context, topic string, data []byte, timeout time.Duration) ([]byte, error) {
	return r.bus.Request(ctx, topic, data, timeout)
}

func (r *Requester[THope]) requestRawAsync(ctx context.Context, topic string, data []byte, timeout time.Duration) ([]byte, error) {
	responses := make(chan []byte)
	gr, ctx := errgroup.WithContext(ctx)
	gr.Go(r.bus.RequestAsync(ctx, topic, data, timeout, responses))
	go func(rsp chan []byte) {
		select {
		case <-ctx.Done():
			close(rsp)
		case rsp <- <-responses:
		}
	}(responses)
	return <-responses, gr.Wait()
}

func (r *Requester[THope]) RequestAsync(ctx context.Context, hope THope, timeout time.Duration) dtos.IFbk {
	fbk := dtos.NewFbk(hope.GetId(), -1, "")
	data, err := convert.Any2Data(hope)
	if err != nil {
		fbk.SetError(err.Error())
		return fbk
	}
	f, err := r.requestRawAsync(ctx, r.Topic, data, timeout)
	if err != nil {
		fbk.SetError(err.Error())
		return fbk
	}
	err = convert.Data2Any(f, &fbk)
	if err != nil {
		fbk.SetError(err.Error())
		return fbk
	}
	return fbk
}

func (r *Requester[THope]) Request(ctx context.Context, hope THope, timeout time.Duration) dtos.IFbk {
	fbk := dtos.NewFbk(hope.GetId(), -1, "")
	data, err := convert.Any2Data(hope)
	if err != nil {
		fbk.SetError(err.Error())
		return fbk
	}
	f, err := r.requestRaw(ctx, r.Topic, data, timeout)
	if err != nil {
		fbk.SetError(err.Error())
		return fbk
	}
	err = convert.Data2Any(f, &fbk)
	if err != nil {
		fbk.SetError(err.Error())
		return fbk
	}
	return fbk
}

func NewRequester[THope dtos.IHope](topic string) (*Requester[THope], error) {
	return newRequester[THope](topic)
}

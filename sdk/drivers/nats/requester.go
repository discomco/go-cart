package nats

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/schema"
	"golang.org/x/sync/errgroup"
	"time"
)

type IRequester[THope contract.IHope] interface {
	comps.IGenRequester[THope]
}

type Requester[THope contract.IHope] struct {
	*comps.Component
	bus   INATSBus
	Topic string
}

const (
	RequesterFmt = "NATS.Requester(%+v)"
)

func (r *Requester[THope]) GetHopeType() contract.HopeType {
	return contract.HopeType(r.Topic)
}

func (r *Requester[THope]) IAmRequester() {}

func newRequester[THope contract.IHope](topic string) (*Requester[THope], error) {
	name := fmt.Sprintf(RequesterFmt, topic)
	ac := comps.NewComponent(schema.Name(name))
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
		bus:       b,
		Component: ac,
		Topic:     topic,
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

func (r *Requester[THope]) GenRequestAsync(ctx context.Context, hope THope, timeout time.Duration) contract.IFbk {
	fbk := contract.NewFbk(hope.GetId(), -1, "")
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

func (r *Requester[THope]) RequestAsync(ctx context.Context, hope contract.IHope, timeout time.Duration) contract.IFbk {
	fbk := contract.NewFbk(hope.GetId(), -1, "")
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

func (r *Requester[THope]) GenRequest(ctx context.Context, hope THope, timeout time.Duration) contract.IFbk {
	fbk := contract.NewFbk(hope.GetId(), -1, "")
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

func (r *Requester[THope]) Request(ctx context.Context, hope contract.IHope, timeout time.Duration) contract.IFbk {
	fbk := contract.NewFbk(hope.GetId(), -1, "")
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

func NewRequester[THope contract.IHope](topic string) (*Requester[THope], error) {
	return newRequester[THope](topic)
}

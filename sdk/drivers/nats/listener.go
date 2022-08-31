package nats

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

type IListener[TFact dtos.IFact, TCmd domain.ICmd] interface {
	features.IGenFactListener[*nats.Msg, TFact]
}

type Listener[TCmd domain.ICmd] struct {
	*features.AppComponent
	Topic         string
	natsBus       INATSBus
	newCmdHandler features.CmdHandlerFtor
	data2Cmd      domain.GenData2CmdFunc[TCmd]
}

func NewListener[TCmd domain.ICmd](
	topic string,
	d2c domain.GenData2CmdFunc[TCmd]) (*Listener[TCmd], error) {
	return newListener[TCmd](topic, d2c)
}

const (
	ListenerFmt = "%+v.NATSListener"
)

func newListener[TCmd domain.ICmd](
	topic string,
	data2Cmd domain.GenData2CmdFunc[TCmd],
) (*Listener[TCmd], error) {
	l := &Listener[TCmd]{
		Topic:    topic,
		data2Cmd: data2Cmd,
	}
	name := fmt.Sprintf(ListenerFmt, topic)
	b := features.NewAppComponent(features.Name(name))
	dig := ioc.SingleIoC()
	var err error
	err = dig.Invoke(func(newBus BusFtor, newCH features.CmdHandlerFtor) {
		l.natsBus, err = newBus()
		l.newCmdHandler = newCH
	})
	if err != nil {
		l.GetLogger().Fatal(err)
		return nil, err
	}
	l.AppComponent = b
	return l, nil
}

func (l *Listener[TCmd]) IAmListener() {}

func (l *Listener[TCmd]) Activate(ctx context.Context) error {
	l.GetLogger().Debugf("%+v ~> activated", l.GetName())
	g, ctx := errgroup.WithContext(ctx)
	g.Go(l.worker(ctx))
	return g.Wait()
}

func (l *Listener[TCmd]) Deactivate(ctx context.Context) error {
	l.GetLogger().Debugf("%+v ~> deactivated", l.GetName())
	l.natsBus.Close()
	return nil
}

func (l *Listener[TCmd]) worker(ctx context.Context) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			{
				err := l.Deactivate(ctx)
				if err != nil {
					return err
				}
				return ctx.Err()
			}
		default:
			logger := l.GetLogger()
			logger.Infof("%+v ~> %+v", l.GetName(), l.natsBus.Connection().Status())
			factChan := make(chan []byte)
			l.natsBus.ListenAsync(ctx, l.Topic, factChan)
			//			_ = l.natsBus.Wait()
			for {
				msg := <-factChan
				var fbk dtos.IFbk
				fbk = dtos.NewFbk("", -1, "")
				cmd, err := l.data2Cmd(msg)
				if err != nil {
					l.handleError(err, fbk)
					continue
				} else {
					cmdHandler := l.newCmdHandler()
					fbk = cmdHandler.Handle(ctx, cmd)
					l.GetLogger().Infof("\n\tFact [%+v]\n\tFbk [%+v]\n", msg, fbk)
				}
			}
		}
	}
}

func (l *Listener[TCmd]) handleError(err error, fbk dtos.IFbk) {
	if err != nil {
		l.GetLogger().Debug(err)
		fbk.SetError(err.Error())
		return
	}
}

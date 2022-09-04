package nats

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

type IListener[TFact contract.IFact, TCmd behavior.ICmd] interface {
	comps.IGenFactListener[*nats.Msg, TFact]
}

type Listener[TCmd behavior.ICmd] struct {
	*comps.Component
	Topic         string
	natsBus       INATSBus
	newCmdHandler comps.CmdHandlerFtor
	data2Cmd      behavior.GenData2CmdFunc[TCmd]
}

func NewListener[TCmd behavior.ICmd](
	topic string,
	d2c behavior.GenData2CmdFunc[TCmd]) (*Listener[TCmd], error) {
	return newListener[TCmd](topic, d2c)
}

const (
	ListenerFmt = "%+v.NATSListener"
)

func newListener[TCmd behavior.ICmd](
	topic string,
	data2Cmd behavior.GenData2CmdFunc[TCmd],
) (*Listener[TCmd], error) {
	l := &Listener[TCmd]{
		Topic:    topic,
		data2Cmd: data2Cmd,
	}
	name := fmt.Sprintf(ListenerFmt, topic)
	b := comps.NewComponent(schema.Name(name))
	dig := ioc.SingleIoC()
	var err error
	err = dig.Invoke(func(newBus BusFtor, newCH comps.CmdHandlerFtor) {
		l.natsBus, err = newBus()
		l.newCmdHandler = newCH
	})
	if err != nil {
		l.GetLogger().Fatal(err)
		return nil, err
	}
	l.Component = b
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
				var fbk contract.IFbk
				fbk = contract.NewFbk("", -1, "")
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

func (l *Listener[TCmd]) handleError(err error, fbk contract.IFbk) {
	if err != nil {
		l.GetLogger().Debug(err)
		fbk.SetError(err.Error())
		return
	}
}

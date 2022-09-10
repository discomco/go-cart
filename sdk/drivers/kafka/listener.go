package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type IListener[TFact contract.IFact, TCmd behavior.ICmd] interface {
	comps.IGenListener[*kafka.Event, TFact]
}

type Listener[TCmd behavior.ICmd] struct {
	*comps.Component
	Topic         string
	consumer      *kafka.Consumer
	newCmdHandler comps.CmdHandlerFtor
	data2Cmd      behavior.GenData2CmdFunc[TCmd]
}

func (l *Listener[TCmd]) IAmListener() {}

func newListener[TCmd behavior.ICmd](
	name schema.Name,
	topic string,
	data2Cmd behavior.GenData2CmdFunc[TCmd],
) (*Listener[TCmd], error) {
	l := &Listener[TCmd]{
		Topic:    topic,
		data2Cmd: data2Cmd,
	}
	b := comps.NewComponent(name)
	dig := ioc.SingleIoC()
	var err error
	err = dig.Invoke(func(newConsumer ConsumerFtor, newCH comps.CmdHandlerFtor) {
		l.newCmdHandler = newCH
		l.consumer, err = newConsumer()
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create consumer: %v", err)
	}
	l.Component = b
	return l, nil
}

func (l *Listener[TCmd]) Activate(ctx context.Context) error {
	l.GetLogger().Debugf("%+v ~> activated", l.GetName())
	g, ctx := errgroup.WithContext(ctx)
	g.Go(l.worker(ctx))
	return g.Wait()
}

func (l *Listener[TCmd]) Deactivate(ctx context.Context) error {
	l.GetLogger().Debugf("%+v ~> deactivated", l.GetName())
	l.consumer.Close()
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
			l.consumer.Subscribe(l.Topic, nil)
			for {
				msg, err := l.consumer.ReadMessage(-1)
				var fbk contract.IFbk
				fbk = contract.NewFbk("", -1, "")
				cmd, err := l.data2Cmd(msg.Value)
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

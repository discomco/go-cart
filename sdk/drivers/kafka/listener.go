package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/discomco/go-cart/core/ioc"
	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/dtos"
	"github.com/discomco/go-cart/features"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type IListener[TFact dtos.IFact, TCmd domain.ICmd] interface {
	features.IGenFactListener[*kafka.Event, TFact]
}

type Listener[TCmd domain.ICmd] struct {
	*features.AppComponent
	Topic         string
	consumer      *kafka.Consumer
	newCmdHandler features.CmdHandlerFtor
	data2Cmd      domain.GenData2CmdFunc[TCmd]
}

func (l *Listener[TCmd]) IAmListener() {}

func newListener[TCmd domain.ICmd](
	name features.Name,
	topic string,
	data2Cmd domain.GenData2CmdFunc[TCmd],
) (*Listener[TCmd], error) {
	l := &Listener[TCmd]{
		Topic:    topic,
		data2Cmd: data2Cmd,
	}
	b := features.NewAppComponent(name)
	dig := ioc.SingleIoC()
	var err error
	err = dig.Invoke(func(newConsumer ConsumerFtor, newCH features.CmdHandlerFtor) {
		l.newCmdHandler = newCH
		l.consumer, err = newConsumer()
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create consumer: %v", err)
	}
	l.AppComponent = b
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
				var fbk dtos.IFbk
				fbk = dtos.NewFbk("", -1, "")
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

func (l *Listener[TCmd]) handleError(err error, fbk dtos.IFbk) {
	if err != nil {
		l.GetLogger().Debug(err)
		fbk.SetError(err.Error())
		return
	}
}

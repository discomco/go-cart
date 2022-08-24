package nats

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/discomco/go-cart/config"
	"github.com/discomco/go-cart/features"
	"github.com/nats-io/nats.go"
	"golang.org/x/sync/errgroup"
)

// INATSBus is the Injector that hides driver specifics from the application.
type INATSBus interface {
	features.IGenBus[*nats.Conn, *nats.Msg]
}

type BusFtor func() (INATSBus, error)

var (
	cMutex     = &sync.Mutex{}
	singleNats INATSBus
)

// SingleNATS returns functor that produces a guaranteed singleton instance of the NATS bus.
func SingleNATS(config config.IAppConfig) BusFtor {
	return func() (INATSBus, error) {
		return singleton(config.GetNATSConfig())
	}
}

// TransientNATS returns functor that produces a transient instance of the NATS bus.
func TransientNATS(config config.IAppConfig) BusFtor {
	return func() (INATSBus, error) {
		return transient(config.GetNATSConfig())
	}
}

func IAmNATSBus(b INATSBus) bool {
	return true
}

type bus struct {
	*features.AppComponent
	conn  *nats.Conn
	mutex *sync.Mutex
	wg    *errgroup.Group
}

func (b *bus) Publish(ctx context.Context, topic string, data []byte) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		b.GetLogger().Infof("Publishing nats Fact [%s]\n", topic)
		err := b.conn.Publish(topic, data)
		if err != nil {
			b.GetLogger().Errorf("Error publishing NATS Topic [%s] Error: %v\n", topic, err)
			return err
		}
		err = b.conn.Flush()
		if err != nil {
			b.GetLogger().Errorf("Error flushing NATS Connection Error: %v\n", err)
			return err
		}
		return nil
	}
}

func (b *bus) Listen(ctx context.Context, topic string, facts chan []byte) {
	if b.wg == nil {
		b.wg, ctx = errgroup.WithContext(ctx)
	}
	b.wg.Go(b.ListenAsync(ctx, topic, facts))
}

func (b *bus) Wait() error {
	return b.wg.Wait()
}

func (b *bus) ListenAsync(ctx context.Context, topic string, facts chan []byte) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			sub, err := b.conn.Subscribe(topic,
				func(msg *nats.Msg) {
					facts <- msg.Data
				})
			if err != nil {
				b.GetLogger().Errorf("Error subscribing to NATS Topic [%s] Error: %v\n", topic, err)
				return err
			}
			err = b.conn.Flush()
			if err != nil {
				b.GetLogger().Errorf("Error flushing NATS Connection Error: %v\n", err)
				return err
			}
			b.GetLogger().Infof("Listening to nats Fact [%s]\n", sub.Subject)
			select {}
		}
	}
}

func (b *bus) Respond(ctx context.Context, topic string, hopes chan *nats.Msg) {
	if b.wg == nil {
		b.wg, ctx = errgroup.WithContext(ctx)
	}
	b.wg.Go(b.RespondAsync(ctx, topic, hopes))
}

func (b *bus) RespondAsync(ctx context.Context, topic string, hopes chan *nats.Msg) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			sub, err := b.conn.Subscribe(topic, func(msg *nats.Msg) {
				hopes <- msg
			})
			if err != nil {
				b.GetLogger().Errorf("Error subscribing [%s] Error: %v\n", topic, err)
				return err
			}
			err = b.conn.Flush()
			if err != nil {
				b.GetLogger().Errorf("Flushing Error: %v\n", err)
				return err
			}
			b.GetLogger().Infof("Responding to [%s]", sub.Subject)
		}
		return nil
	}
}

func (b *bus) Request(ctx context.Context, topic string, data []byte, timeout time.Duration) ([]byte, error) {
	msg, err := b.conn.Request(topic, data, timeout)
	if err != nil {
		b.GetLogger().Errorf("Error requesting [%s] Error: %v\n", topic, err)
		return nil, err
	}
	return msg.Data, nil
}

func (b *bus) RequestAsync(ctx context.Context, topic string, data []byte, timeout time.Duration, responses chan []byte) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := b.conn.Request(topic, data, timeout)
			if err != nil {
				b.GetLogger().Errorf("Error requesting [%s] Error: %v\n", topic, err)
				return err
			}
			responses <- msg.Data
		}
		return nil
	}
}

func (b *bus) Connection() *nats.Conn {
	return b.conn
}

func singleton(cfg config.INATSConfig) (INATSBus, error) {
	if singleNats == nil {
		cMutex.Lock()
		defer cMutex.Unlock()
		s, err := transient(cfg)
		if err != nil {
			return nil, err
		}
		singleNats = s
	}
	return singleNats, nil
}

const (
	BusFmt = "[%v].NATS.New"
)

func transient(cfg config.INATSConfig) (INATSBus, error) {
	name := fmt.Sprintf(BusFmt, cfg.GetUrl())
	b := &bus{
		AppComponent: features.NewAppComponent(features.Name(name)),
		wg:           &errgroup.Group{},
	}
	conn, err := nats.Connect(
		cfg.GetUrl(),
		nats.UserInfo(cfg.GetUser(), cfg.GetPwd()))
	if err != nil {
		b.GetLogger().Errorf("Error connecting to [%s] Error: %v\n", cfg.GetUrl(), err)
		return nil, err
	}
	b.conn = conn
	return b, nil
}

func (b *bus) Close() {
	b.conn.Close()
}

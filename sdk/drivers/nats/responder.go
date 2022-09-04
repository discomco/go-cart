package nats

import (
	"encoding/json"
	"fmt"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/comps"
	"github.com/discomco/go-cart/sdk/contract"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/nats-io/nats.go"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"sync"
)

type IResponder[THope contract.IHope, TCmd behavior.ICmd] interface {
	comps.IResponder
}

type Responder[THope contract.IHope, TCmd behavior.ICmd] struct {
	*comps.Component
	Topic         string
	natsBus       INATSBus
	newCmdHandler comps.CmdHandlerFtor
	hope2Cmd      behavior.Hope2CmdFunc[THope, TCmd]
	mMutex        *sync.Mutex
}

func (r *Responder[THope, TCmd]) GetHopeType() contract.HopeType {
	return contract.HopeType(r.Topic)
}

func (r *Responder[THope, TCmd]) IAmHopeResponder() {}

//Deactivate is called when the component is deactivated
func (r *Responder[THope, TCmd]) Deactivate(ctx context.Context) error {
	r.GetLogger().Infof("%+v deactivated", r.GetName())
	r.natsBus.Close()
	return nil
}

//ResponderFtor is a generic functor that is discriminated by the feature's specific IHope and ICmd injectors.
func ResponderFtor[THope contract.IHope, TCmd behavior.ICmd](
	topic string,
	feature features.ISpoke,
	hope2Cmd behavior.Hope2CmdFunc[THope, TCmd],
) comps.GenResponderFtor[THope] {
	return func() comps.IGenResponder[THope] {
		r, err := newResponder[THope, TCmd](
			topic,
			hope2Cmd)
		if err != nil {
			panic(err)
		}
		feature.Inject(r)
		return r
	}
}

func NewResponder[THope contract.IHope, TCmd behavior.ICmd](
	topic string,
	h2c behavior.Hope2CmdFunc[THope, TCmd]) (IResponder[THope, TCmd], error) {
	return newResponder[THope, TCmd](topic, h2c)
}

const (
	ResponderFmt = "[%+v].NATSResponder"
)

func newResponder[THope contract.IHope, TCmd behavior.ICmd](
	topic string,
	hope2Cmd behavior.Hope2CmdFunc[THope, TCmd],
) (*Responder[THope, TCmd], error) {
	name := fmt.Sprintf(ResponderFmt, topic)
	r := &Responder[THope, TCmd]{
		Topic:  topic,
		mMutex: &sync.Mutex{},
	}
	c := comps.NewComponent(schema.Name(name))
	dig := ioc.SingleIoC()
	var err error
	dig.Invoke(func(newBus BusFtor, newCH comps.CmdHandlerFtor) {
		r.natsBus, err = newBus()
		r.newCmdHandler = newCH
	})
	if err != nil {
		c.GetLogger().Fatal(err)
		return nil, err
	}
	r.hope2Cmd = hope2Cmd
	r.Component = c
	return r, nil
}

func (r *Responder[THope, TCmd]) Activate(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(r.worker(ctx))
	r.GetLogger().Infof("%+v activated", r.GetName())
	return g.Wait()
}

func (r *Responder[THope, TCmd]) mapHope2Cmd(hope *contract.Dto) (TCmd, error) {
	r.mMutex.Lock()
	defer r.mMutex.Unlock()
	return r.hope2Cmd(hope)
}

func (r *Responder[THope, TCmd]) worker(ctx context.Context) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			{
				err := r.Deactivate(ctx)
				if err != nil {
					return err
				}
				return ctx.Err()
			}
		default:
			logger := r.GetLogger()
			logger.Infof("%+v ~> %v", r.GetName(), r.natsBus.Connection().Status())
			hopeChan := make(chan *nats.Msg)
			//g, ctx := errgroup.WithContext(ctx)
			r.natsBus.Respond(ctx, r.Topic, hopeChan)
			//			_ = r.natsBus.Wait()
			for {
				msg := <-hopeChan
				var fbk contract.IFbk
				fbk = contract.NewFbk("", -1, "")

				// TODO: Remove Debugging Code
				logger.Debugf("[%+v] received %+v", r.GetName(), string(msg.Data))

				var dto contract.Dto
				err := convert.Data2Any(msg.Data, &dto)

				// TODO: Remove Debugging Code
				logger.Debugf("[%+v] converted %+v to %+v", r.GetName(), string(msg.Data), dto)

				if err != nil {
					r.handleError(err, fbk, msg)
					continue
				}
				cmd, err := r.mapHope2Cmd(&dto)
				if err != nil {
					r.handleError(err, fbk, msg)
					continue
				}
				cmdHandler := r.newCmdHandler()
				fbk = cmdHandler.Handle(ctx, cmd)
				rsp, err := json.Marshal(fbk)
				if err != nil {
					r.handleError(err, fbk, msg)
					continue
				}
				err = msg.Respond(rsp)
				if err != nil {
					r.GetLogger().Debug(err)
					continue
				}
			}
		}
	}
}

func (r *Responder[THope, TCmd]) handleError(err error, fbk contract.IFbk, msg *nats.Msg) {
	fbk.SetError(err.Error())
	rsp, err := json.Marshal(fbk)
	if err != nil {
		r.GetLogger().Debug(err)
		return
	}
	err = msg.Respond(rsp)
	if err != nil {
		r.GetLogger().Debug(err)
		return
	}
}

package nats

import (
	"encoding/json"
	"fmt"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/core/utils/convert"
	"github.com/discomco/go-cart/sdk/domain"
	"github.com/discomco/go-cart/sdk/dtos"
	"github.com/discomco/go-cart/sdk/features"
	"github.com/nats-io/nats.go"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"sync"
)

type INATSResponder[THope dtos.IHope, TCmd domain.ICmd] interface {
	features.IHopeResponder
}

type Responder[THope dtos.IHope, TCmd domain.ICmd] struct {
	*features.AppComponent
	Topic         string
	natsBus       INATSBus
	newCmdHandler features.CmdHandlerFtor
	hope2Cmd      domain.Hope2CmdFunc[THope, TCmd]
	mMutex        *sync.Mutex
}

func (r *Responder[THope, TCmd]) GetHopeType() dtos.HopeType {
	return dtos.HopeType(r.Topic)
}

func (r *Responder[THope, TCmd]) IAmHopeResponder() {}

//Deactivate is called when the component is deactivated
func (r *Responder[THope, TCmd]) Deactivate(ctx context.Context) error {
	r.GetLogger().Infof("%+v deactivated", r.GetName())
	r.natsBus.Close()
	return nil
}

//ResponderFtor is a generic functor that is discriminated by the feature's specific IHope and ICmd injectors.
func ResponderFtor[THope dtos.IHope, TCmd domain.ICmd](
	topic string,
	feature features.IFeature,
	hope2Cmd domain.Hope2CmdFunc[THope, TCmd],
) features.GenResponderFtor[THope] {
	return func() features.IGenResponder[THope] {
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

func NewResponder[THope dtos.IHope, TCmd domain.ICmd](
	topic string,
	h2c domain.Hope2CmdFunc[THope, TCmd]) (INATSResponder[THope, TCmd], error) {
	return newResponder[THope, TCmd](topic, h2c)
}

const (
	ResponderFmt = "[%+v].NATSResponder"
)

func newResponder[THope dtos.IHope, TCmd domain.ICmd](
	topic string,
	hope2Cmd domain.Hope2CmdFunc[THope, TCmd],
) (*Responder[THope, TCmd], error) {
	name := fmt.Sprintf(ResponderFmt, topic)
	r := &Responder[THope, TCmd]{
		Topic:  topic,
		mMutex: &sync.Mutex{},
	}
	c := features.NewAppComponent(features.Name(name))
	dig := ioc.SingleIoC()
	var err error
	dig.Invoke(func(newBus BusFtor, newCH features.CmdHandlerFtor) {
		r.natsBus, err = newBus()
		r.newCmdHandler = newCH
	})
	if err != nil {
		c.GetLogger().Fatal(err)
		return nil, err
	}
	r.hope2Cmd = hope2Cmd
	r.AppComponent = c
	return r, nil
}

func (r *Responder[THope, TCmd]) Activate(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(r.worker(ctx))
	r.GetLogger().Infof("%+v activated", r.GetName())
	return g.Wait()
}

func (r *Responder[THope, TCmd]) mapHope2Cmd(hope *dtos.Dto) (TCmd, error) {
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
				var fbk dtos.IFbk
				fbk = dtos.NewFbk("", -1, "")

				// TODO: Remove Debugging Code
				logger.Debugf("[%+v] received %v", r.GetName(), msg)

				var dto dtos.Dto
				err := convert.Data2Any(msg.Data, &dto)

				// TODO: Remove Debugging Code
				logger.Debugf("[%+v] converted %v to %v", r.GetName(), msg.Data, dto)

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

func (r *Responder[THope, TCmd]) handleError(err error, fbk dtos.IFbk, msg *nats.Msg) {
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

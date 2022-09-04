package reactors

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"sync"
)

type IMsgReactor interface {
	IComponent
	IActivate
	IDeactivate
	//IAmMsgHandler()
	//GetMsgType() MsgType
	//UnsubscribeAll(when OnMsgFunc) map[string]error
	//Deactivate(ctx context.Context) error
	//Unsubscribe(topic string, fn OnMsgFunc) error
	//React(ctx context.Context, msgmodel.IMsg) error
	//Activate(ctx context.Context) error
	//SubscribeAll(ctx context.Context, when OnMsgFunc, transactional bool)
	//Subscribe(ctx context.Context, topic string, when OnMsgFunc, transactional bool)
	//SubscribeAllAsync(msgs chanmodel.IMsg, transactional bool) map[string]error
	//SubscribeAsync(msgs chanmodel.IMsg, transactional bool) error
}

type IGenMsgReactor[TMsg schema.IMsg] interface {
	IMsgReactor
	GenWhen(ctx context.Context, msg TMsg)
}

type OnMsgFunc func(ctx context.Context, msg schema.IMsg) error

type MsgReactorFtor func() IMsgReactor
type GenMsgReactorFtor[TMsg schema.IMsg] func() IGenMsgReactor[TMsg]

type MsgReactor struct {
	*Component
	mediator  mediator.IMediator
	msgType   schema.MsgType
	react     OnMsgFunc
	whenMutex *sync.Mutex
}

const MsgReactorFmt = "%+v.MsgReactor"
const AllTopics = "*"

func NewMsgReactor(
	msgType schema.MsgType,
	react OnMsgFunc,
) *MsgReactor {
	name := fmt.Sprintf(MsgReactorFmt, msgType)
	base := NewComponent(schema.Name(name))
	result := &MsgReactor{
		Component: base,
		msgType:   msgType,
		react:     react,
		whenMutex: &sync.Mutex{},
	}
	return result
}

func (h *MsgReactor) Deactivate(ctx context.Context) error {
	if h.msgType == AllTopics {
		h.UnsubscribeAll(h.When)
		return nil
	}
	err := h.unsub(string(h.msgType), h.When)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	return nil

}

func (h *MsgReactor) Unsubscribe(topic string, fn OnMsgFunc) error {
	return h.unsub(topic, fn)
}

func (h *MsgReactor) unsub(topic string, fn interface{}) error {
	err := h.GetMediator().Unregister(topic, fn)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	h.GetLogger().Infof("[%s] unlinked [%+v]", h.GetName(), h.GetMsgType())
	return nil
}

func (h *MsgReactor) GetMsgType() schema.MsgType {
	return h.msgType
}

func (h *MsgReactor) When(ctx context.Context, msg schema.IMsg) error {
	if h.react == nil {
		return nil
	}
	h.whenMutex.Lock()
	defer h.whenMutex.Unlock()
	h.GetLogger().Infof("[%+v] received [%+v]", h.GetName(), msg)
	return h.react(ctx, msg)
}

func (h *MsgReactor) Activate(ctx context.Context) error {
	if h.msgType == AllTopics {
		h.SubscribeAll(ctx, h.When, true)
		return nil
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, string(h.msgType), h.When, true))
	return wg.Wait()
}

func (h *MsgReactor) subWorker(ctx context.Context, topic string, fn interface{}, transactional bool) func() error {
	return func() error {
		select {
		case <-ctx.Done():
			h.Deactivate(ctx)
			return ctx.Err()
		default:
			err := h.GetMediator().RegisterAsync(topic, fn, transactional)
			if err != nil {
				h.GetLogger().Error(err)
				return errors.Wrap(err, "failed to register with mediator")
			}
			h.GetMediator().WaitAsync()
			h.GetLogger().Infof("[%+v] links [%+v]", h.GetName(), h.msgType)
			return nil
		}
	}
}

func (h *MsgReactor) SubscribeAll(ctx context.Context, when OnMsgFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		wg.Go(h.subWorker(ctx, topic, when, transactional))
	}
	return wg.Wait()
}

func (h *MsgReactor) Subscribe(ctx context.Context, topic string, when OnMsgFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, topic, when, transactional))
	return wg.Wait()
}

func (h *MsgReactor) SubscribeAllAsync(msgs chan schema.IMsg, transactional bool) map[string]error {
	res := make(map[string]error)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		err := h.subAsync(topic, msgs, transactional)
		if err != nil {
			h.GetLogger().Errorf("[%+v] failed to link [%s]. error: %v", h.GetName(), topic, err)
			res[topic] = err
		}
	}
	return res
}

func (h *MsgReactor) subAsync(topic string, msgs chan schema.IMsg, transactional bool) error {
	err := h.GetMediator().RegisterAsync(topic, func(msg schema.IMsg) {
		msgs <- msg
	}, transactional)
	if err != nil {
		h.GetLogger().Fatal(err)
		return err
	}
	h.GetMediator().WaitAsync()
	return nil
}

func (h *MsgReactor) SubscribeAsync(msgs chan schema.IMsg, transactional bool) error {
	return h.subAsync(string(h.msgType), msgs, transactional)
}

func (h *MsgReactor) UnsubscribeAll(when OnMsgFunc) map[string]error {
	errs := make(map[string]error, 0)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		err := h.unsub(topic, when)
		if err != nil {
			errs[topic] = err
		}
	}
	return errs
}

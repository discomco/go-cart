package comps

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"sync"
)

type IMsgReaction interface {
	IReaction
	//IComponent
	//IActivate
	//IDeactivate
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

type IGenMsgReaction[TMsg schema.IMsg] interface {
	IMsgReaction
	GenWhen(ctx context.Context, msg TMsg)
}

type OnMsgFunc func(ctx context.Context, msg schema.IMsg) error

type MsgReactionFtor func() IMsgReaction
type GenMsgReactionFtor[TMsg schema.IMsg] func() IGenMsgReaction[TMsg]

type MsgReaction struct {
	*Component
	mediator  mediator.IMediator
	msgType   schema.MsgType
	react     OnMsgFunc
	whenMutex *sync.Mutex
}

const MsgReactionFmt = "%+v.MsgReaction"
const AllTopics = "*"

func NewMsgReaction(
	msgType schema.MsgType,
	react OnMsgFunc,
) *MsgReaction {
	name := fmt.Sprintf(MsgReactionFmt, msgType)
	base := NewComponent(schema.Name(name))
	result := &MsgReaction{
		Component: base,
		msgType:   msgType,
		react:     react,
		whenMutex: &sync.Mutex{},
	}
	return result
}

func (h *MsgReaction) Deactivate(ctx context.Context) error {
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

func (h *MsgReaction) Unsubscribe(topic string, fn OnMsgFunc) error {
	return h.unsub(topic, fn)
}

func (h *MsgReaction) unsub(topic string, fn interface{}) error {
	err := h.GetMediator().Unregister(topic, fn)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	h.GetLogger().Infof("[%s] unlinked [%+v]", h.GetName(), h.GetMsgType())
	return nil
}

func (h *MsgReaction) GetMsgType() schema.MsgType {
	return h.msgType
}

func (h *MsgReaction) When(ctx context.Context, msg schema.IMsg) error {
	if h.react == nil {
		return nil
	}
	h.whenMutex.Lock()
	defer h.whenMutex.Unlock()
	h.GetLogger().Infof("[%+v] received [%+v]", h.GetName(), msg)
	return h.react(ctx, msg)
}

func (h *MsgReaction) Activate(ctx context.Context) error {
	if h.msgType == AllTopics {
		h.SubscribeAll(ctx, h.When, true)
		return nil
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, string(h.msgType), h.When, true))
	return wg.Wait()
}

func (h *MsgReaction) subWorker(ctx context.Context, topic string, fn interface{}, transactional bool) func() error {
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

func (h *MsgReaction) SubscribeAll(ctx context.Context, when OnMsgFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		wg.Go(h.subWorker(ctx, topic, when, transactional))
	}
	return wg.Wait()
}

func (h *MsgReaction) Subscribe(ctx context.Context, topic string, when OnMsgFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, topic, when, transactional))
	return wg.Wait()
}

func (h *MsgReaction) SubscribeAllAsync(msgs chan schema.IMsg, transactional bool) map[string]error {
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

func (h *MsgReaction) subAsync(topic string, msgs chan schema.IMsg, transactional bool) error {
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

func (h *MsgReaction) SubscribeAsync(msgs chan schema.IMsg, transactional bool) error {
	return h.subAsync(string(h.msgType), msgs, transactional)
}

func (h *MsgReaction) UnsubscribeAll(when OnMsgFunc) map[string]error {
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

package features

import (
	"context"
	"fmt"
	"github.com/discomco/go-cart/sdk/core/mediator"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"sync"
)

type IMsgHandler interface {
	IComponent
	IActivate
	IDeactivate
	//IAmMsgHandler()
	//GetMsgType() MsgType
	//UnsubscribeAll(when OnMsgFunc) map[string]error
	//Deactivate(ctx context.Context) error
	//Unsubscribe(topic string, fn OnMsgFunc) error
	//When(ctx context.Context, msg IMsg) error
	//Activate(ctx context.Context) error
	//SubscribeAll(ctx context.Context, when OnMsgFunc, transactional bool)
	//Subscribe(ctx context.Context, topic string, when OnMsgFunc, transactional bool)
	//SubscribeAllAsync(msgs chan IMsg, transactional bool) map[string]error
	//SubscribeAsync(msgs chan IMsg, transactional bool) error
}

type IGenMsgHandler[TMsg IMsg] interface {
	IMsgHandler
	GenWhen(ctx context.Context, msg TMsg)
}

type OnMsgFunc func(ctx context.Context, msg IMsg) error

type MsgHandlerFtor func() IMsgHandler
type GenMsgHandlerFtor[TMsg IMsg] func() IGenMsgHandler[TMsg]

type MsgHandler struct {
	*AppComponent
	mediator  mediator.IMediator
	msgType   MsgType
	onMsg     OnMsgFunc
	whenMutex *sync.Mutex
}

const MsgHandlerFmt = "%+v.MsgHandler"
const AllTopics = "*"

func NewMsgHandler(
	msgType MsgType,
	onMsg OnMsgFunc,
) *MsgHandler {
	name := fmt.Sprintf(MsgHandlerFmt, msgType)
	base := NewAppComponent(Name(name))
	result := &MsgHandler{
		AppComponent: base,
		msgType:      msgType,
		onMsg:        onMsg,
		whenMutex:    &sync.Mutex{},
	}
	return result
}

func (h *MsgHandler) Deactivate(ctx context.Context) error {
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

func (h *MsgHandler) Unsubscribe(topic string, fn OnMsgFunc) error {
	return h.unsub(topic, fn)
}

func (h *MsgHandler) unsub(topic string, fn interface{}) error {
	err := h.GetMediator().Unregister(topic, fn)
	if err != nil {
		h.GetLogger().Error(err)
		return err
	}
	h.GetLogger().Infof("[%s] unlinked [%+v]", h.GetName(), h.GetMsgType())
	return nil
}

func (h *MsgHandler) GetMsgType() MsgType {
	return h.msgType
}

func (h *MsgHandler) When(ctx context.Context, msg IMsg) error {
	if h.onMsg == nil {
		return nil
	}
	h.whenMutex.Lock()
	defer h.whenMutex.Unlock()
	h.GetLogger().Infof("[%+v] received [%+v]", h.GetName(), msg)
	return h.onMsg(ctx, msg)
}

func (h *MsgHandler) Activate(ctx context.Context) error {
	if h.msgType == AllTopics {
		h.SubscribeAll(ctx, h.When, true)
		return nil
	}
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, string(h.msgType), h.When, true))
	return wg.Wait()
}

func (h *MsgHandler) subWorker(ctx context.Context, topic string, fn interface{}, transactional bool) func() error {
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

func (h *MsgHandler) SubscribeAll(ctx context.Context, when OnMsgFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	topics := h.GetMediator().KnownTopics()
	for topic := range topics {
		wg.Go(h.subWorker(ctx, topic, when, transactional))
	}
	return wg.Wait()
}

func (h *MsgHandler) Subscribe(ctx context.Context, topic string, when OnMsgFunc, transactional bool) error {
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(h.subWorker(ctx, topic, when, transactional))
	return wg.Wait()
}

func (h *MsgHandler) SubscribeAllAsync(msgs chan IMsg, transactional bool) map[string]error {
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

func (h *MsgHandler) subAsync(topic string, msgs chan IMsg, transactional bool) error {
	err := h.GetMediator().RegisterAsync(topic, func(msg IMsg) {
		msgs <- msg
	}, transactional)
	if err != nil {
		h.GetLogger().Fatal(err)
		return err
	}
	h.GetMediator().WaitAsync()
	return nil
}

func (h *MsgHandler) SubscribeAsync(msgs chan IMsg, transactional bool) error {
	return h.subAsync(string(h.msgType), msgs, transactional)
}

func (h *MsgHandler) UnsubscribeAll(when OnMsgFunc) map[string]error {
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

package features

import (
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/core"
	"github.com/discomco/go-cart/core/builder"
	"github.com/discomco/go-cart/core/ioc"
	"github.com/discomco/go-cart/domain"
	"github.com/discomco/go-cart/drivers/convert"
	"github.com/discomco/go-cart/model"
	"github.com/discomco/go-cart/test"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"reflect"
	"sync"
)

const (
	CfgPath           = "../config/config.yaml"
	MY_PAYLOAD_PREFIX = "my"
)

var (
	testEnv ioc.IDig
)

func init() {
	testEnv = buildTestEnv()
}

func buildTestEnv() ioc.IDig {
	dig := builder.InjectCoLoMed(CfgPath)
	return dig.Inject(dig,
		NewMyEvtHandler,
	)
}

type iMyEvt interface {
	domain.IEvt
}

type MyPayload struct {
	ID     *core.Identity
	Brand  string
	Model  string
	Status int
}

func newMyPayload(id string) (model.IPayload, error) {
	ID, err := core.NewIdentityFrom(MY_PAYLOAD_PREFIX, id)
	if err != nil {
		err = errors.Wrap(err, "newMyPayload")
		return nil, err
	}
	return &MyPayload{
		ID:     ID,
		Brand:  "Toyota",
		Model:  "Yaris",
		Status: 42,
	}, nil
}

func newMyTestEvt() (iMyEvt, error) {
	// AND WE CREATE AN INITIALIZE EVENT
	eid, _ := uuid.NewV4()
	pl, err := newMyPayload(eid.String())
	if err != nil {
		err = errors.Wrap(err, "newMyPayload")
		return nil, err
	}
	d, _ := json.Marshal(pl)
	evtData := &esdb.EventData{
		EventID:     eid,
		EventType:   MY_EVT_TOPIC,
		ContentType: 0,
		Data:        d,
		Metadata:    nil,
	}
	evt := convert.EventData2Evt(*evtData).(*domain.Event)
	id, err := core.NewIdentityFrom(MY_PAYLOAD_PREFIX, test.TEST_UUID)
	if err != nil {
		err := errors.Wrap(err, "newMyTestEvt")
		return nil, err
	}
	evt.SetAggregateId(id.Id())
	return evt, err
}

type iMyEvtHandler interface {
	IGenMediatorSubscriber[iMyEvt]
}

type myGenEvtHandler struct {
	*GenEvtHandler[iMyEvt]
}

func (h *myGenEvtHandler) handleEvt(ctx context.Context, evt domain.IEvt) error {
	h.GetLogger().Info(reflect.TypeOf(evt))
	return nil
}

var hMutex = &sync.Mutex{}

func newMyGenHandler() iMyEvtHandler {
	h := &myGenEvtHandler{}
	base := newGenEvtHandler[iMyEvt](MY_EVT_TOPIC, h.handleEvt)
	h.GenEvtHandler = base
	return h
}

const (
	MY_EVT_TOPIC = "my-evt-topic"
)

func NewMyEvtHandler() GenEvtHandlerFtor[iMyEvt] {
	return func() IGenEvtHandler[iMyEvt] {
		return newMyGenHandler()
	}
}

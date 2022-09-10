package comps

import (
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/v2/esdb"
	"github.com/discomco/go-cart/sdk/behavior"
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"
	"github.com/discomco/go-cart/sdk/drivers/convert"
	"github.com/discomco/go-cart/sdk/schema"
	"github.com/discomco/go-cart/sdk/test"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"reflect"
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
	behavior.IEvt
}

type MyPayload struct {
	ID     *schema.Identity
	Brand  string
	Model  string
	Status int
}

func newMyPayload(id string) (schema.IPayload, error) {
	ID, err := schema.NewIdentityFrom(MY_PAYLOAD_PREFIX, id)
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
	evt := convert.EventData2Evt(*evtData).(*behavior.Event)
	id, err := schema.NewIdentityFrom(MY_PAYLOAD_PREFIX, test.TEST_UUID)
	if err != nil {
		err := errors.Wrap(err, "newMyTestEvt")
		return nil, err
	}
	evt.SetAggregateId(id.Id())
	return evt, err
}

type iMyEvtHandler interface {
	IGenMediatorReaction[iMyEvt]
}

type myGenEvtHandler struct {
	*GenEvtReactor[iMyEvt]
}

func (h *myGenEvtHandler) handleEvt(ctx context.Context, evt behavior.IEvt) error {
	h.GetLogger().Info(reflect.TypeOf(evt))
	return nil
}

func newMyGenHandler() iMyEvtHandler {
	h := &myGenEvtHandler{}
	base := newGenEvtHandler[iMyEvt](MY_EVT_TOPIC, h.handleEvt)
	h.GenEvtReactor = base
	return h
}

const (
	MY_EVT_TOPIC = "my-evt-topic"
)

func NewMyEvtHandler() GenEvtReactionFtor[iMyEvt] {
	return func() IGenEvtReaction[iMyEvt] {
		return newMyGenHandler()
	}
}

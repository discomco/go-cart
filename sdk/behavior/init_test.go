package behavior

import (
	"github.com/discomco/go-cart/sdk/core/builder"
	"github.com/discomco/go-cart/sdk/core/ioc"

	"github.com/discomco/go-cart/sdk/schema"
)

var testEnv ioc.IDig

func init() {
	testEnv = newTestEnv()
}

func newTestEnv() ioc.IDig {
	ioc := builder.InjectCoLoMed(CfgPath)
	return ioc.Inject(ioc,
		ABehaviorFtor,
		ABehaviorBuilder)
}

type anApply struct {
	*ApplyEvt
}

func (a *anApply) fApply(schema schema.ISchema, evt IEvt) error {
	return nil
}

type IAnApplyEvt interface {
	IApplyEvt
}

func newAnApply() IAnApplyEvt {
	a := &anApply{}
	b := NewApplyEvt(A_EVT_TOPIC, a.fApply)
	a.ApplyEvt = b
	return a
}

func AnApply() IBehaviorPlugin {
	return newAnApply()
}

type IAnEvt interface {
	IEvt
}

func NewAnEvt(aggregate IBehavior) IAnEvt {
	return NewEvt(aggregate, A_EVT_TOPIC)
}

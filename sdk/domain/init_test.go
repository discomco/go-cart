package domain

import (
	"github.com/discomco/go-cart/core/builder"
	"github.com/discomco/go-cart/core/ioc"

	"github.com/discomco/go-cart/model"
)

var testEnv ioc.IDig

func init() {
	testEnv = newTestEnv()
}

func newTestEnv() ioc.IDig {
	ioc := builder.InjectCoLoMed(CfgPath)
	return ioc.Inject(ioc,
		AnAggFtor,
		AnAggBuilder)
}

type anApply struct {
	*ApplyEvt
}

func (a *anApply) applyEvt(evt IEvt, state model.IWriteModel) error {
	return nil
}

type IAnApplyEvt interface {
	IApplyEvt
}

func newAnApply() IAnApplyEvt {
	a := &anApply{}
	b := NewApplyEvt(A_EVT_TOPIC, a.applyEvt)
	a.ApplyEvt = b
	return a
}

func AnApply() IAggPlugin {
	return newAnApply()
}

type IAnEvt interface {
	IEvt
}

func NewAnEvt(aggregate IAggregate) IAnEvt {
	return NewEvt(aggregate, A_EVT_TOPIC)
}

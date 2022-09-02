package domain

import (
	"github.com/discomco/go-cart/robby/execute-game/-shared/model"
	domain2 "github.com/discomco/go-cart/robby/execute-game/modules/initialize/domain"
	sdk_domain "github.com/discomco/go-cart/sdk/domain"
)

type IAggregate interface {
	sdk_domain.IAggregate
}

type AggFtor sdk_domain.GenAggFtor[model.Root]

func NewAgg(newRoot model.DocFtor) AggFtor {
	return func() sdk_domain.IAggregate {
		return sdk_domain.NewAggregate(domain2.EVT_TOPIC, newRoot())
	}
}

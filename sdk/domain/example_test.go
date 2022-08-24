package domain

const (
	A_PREFIX                  = "a"
	CfgPath                   = "../config/config.yaml"
	anAggType   AggregateType = A_PREFIX
	A_CMD_TOPIC               = "a_cmd_topic"
	A_EVT_TOPIC               = "a_event_topic"
)

type IATryCmd interface {
	ITryCmd
}

func AnAggBuilder(ftor AggFtor) AggBuilder {
	return func() IAggregate {
		return ftor()
	}
}

func AnAggFtor() AggFtor {
	return func() IAggregate {
		return NewAggregate(anAggType, newWriteModel())
	}
}

type AWriteModel struct {
	status int
}

func (c *AWriteModel) GetStatus() int {
	return c.status
}
func newWriteModel() *AWriteModel {
	return &AWriteModel{}
}

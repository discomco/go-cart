package behavior

const (
	A_PREFIX                 = "a"
	CfgPath                  = "../config/config.yaml"
	anAggType   BehaviorType = A_PREFIX
	A_CMD_TOPIC              = "a_cmd_topic"
	A_EVT_TOPIC              = "a_event_topic"
)

type IATryCmd interface {
	ITryCmd
}

func AnAggBuilder(ftor BehaviorFtor) BehaviorBuilder {
	return func() IBehavior {
		return ftor()
	}
}

func AnAggFtor() BehaviorFtor {
	return func() IBehavior {
		return NewBehavior(anAggType, newWriteModel())
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

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

func ABehaviorBuilder(ftor BehaviorFtor) BehaviorBuilder {
	return func() IBehavior {
		return ftor()
	}
}

func ABehaviorFtor() BehaviorFtor {
	return func() IBehavior {
		return NewBehavior(anAggType, newWriteModel())
	}
}

type ASchema struct {
	status int
}

func (c *ASchema) GetStatus() int {
	return c.status
}
func newWriteModel() *ASchema {
	return &ASchema{}
}

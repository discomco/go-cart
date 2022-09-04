package behavior

type Topic string
type MediatorTopic string

type ITopic interface {
	Topic() Topic
}

func ImplementsITopic(topic ITopic) bool {
	return true
}

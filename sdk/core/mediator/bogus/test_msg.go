package bogus

// TestMsg
type TestMsg struct {
	Content string
}

func NewTestMsg(content string) *TestMsg {
	return &TestMsg{
		Content: content,
	}
}

func (msg *TestMsg) Topic() string {
	return TestTopic
}

// QuitMsg
type QuitMsg struct{}

func NewQuitMsg() *QuitMsg {
	return &QuitMsg{}
}

func (q *QuitMsg) Topic() string {
	return QuitTopic
}

package message

type Message interface {
	GetText() string
}

// Implements message interface
type BaseMessage struct {
	Text string `json:"text"`
}

func (m *BaseMessage) GetText() string {
	return m.Text
}

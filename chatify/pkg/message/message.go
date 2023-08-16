package message

type Message interface {
	GetUsername() string
}

// Implements message interface
type BaseMessage struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (m *BaseMessage) GetUsername() string {
	return m.Username
}

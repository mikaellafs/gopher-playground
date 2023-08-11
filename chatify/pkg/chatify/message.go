package chatify

type message interface {
	GetUsername() string
}

type BaseMessage struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (m *BaseMessage) GetUsername() string {
	return m.Username
}

package chatify

import "encoding/json"

type Message struct {
	Username string
	Text     string
}

func (m Message) ToBytes() []byte {
	data, _ := json.Marshal(m)
	return data
}

func NewMessageFrom(data []byte) (Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)

	return msg, err
}

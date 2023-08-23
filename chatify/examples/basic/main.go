package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopher-playground/chatify/pkg/chatify"
	"gopher-playground/chatify/pkg/message"
	"time"
)

type CustomMessage struct {
	message.BaseMessage

	Username string `json:"username"`
	Location string `json:"location"`
}

func main() {
	server := chatify.NewServer(
		chatify.WithPort(8000),
		chatify.WithServerPath("/chat"),
	)
	g := server.NewGroup(
		chatify.WithGroupPath("/everyone"),
		chatify.WithMessageFormat(parseMessage),
	)

	go server.Run()

	for {
		<-time.After(15 * time.Second)
		fmt.Println("Total clients:", g.TotalClients())
	}
}

func parseMessage(data []byte) (message.Message, error) {
	// Marshal data
	var msg CustomMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Username == "" {
		return nil, errors.New("missing username")
	}

	// Check location
	if msg.Location != "Brasil" {
		return nil, errors.New("invalid location")
	}

	return message.Message(&msg), nil
}

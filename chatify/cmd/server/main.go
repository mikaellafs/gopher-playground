package main

import (
	"fmt"
	"gopher-playground/chatify/pkg/chatify"
	"gopher-playground/chatify/pkg/message"
	"time"
)

type TestMessageStore struct {
}

func (s *TestMessageStore) SaveMessage(m message.Message) error {
	fmt.Println("Saving message...", m.GetUsername())
	return nil
}

func main() {
	server := chatify.NewServer(
		chatify.WithMessageStore(&TestMessageStore{}),
	)
	go server.Run()

	for {
		<-time.After(15 * time.Second)
		fmt.Println("Total clients:", server.TotalClients())
	}
}

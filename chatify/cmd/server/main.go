package main

import (
	"fmt"
	"gopher-playground/chatify/pkg/chatify"
	"gopher-playground/chatify/pkg/message"
	"time"

	"github.com/gin-gonic/gin"
)

type TestMessageStore struct {
}

func (s *TestMessageStore) SaveMessage(m message.Message) error {
	fmt.Println("Saving message...", m.GetText())
	return nil
}

func main() {
	server := chatify.NewServer(
		chatify.WithServerMiddleware(func(c *gin.Context) {
			fmt.Println("Ihuuu")
		}),
	)
	g := server.NewGroup(
		chatify.WithMessageStore(&TestMessageStore{}),
	)
	go server.Run()

	for {
		<-time.After(15 * time.Second)
		fmt.Println("Total clients:", g.TotalClients())
	}
}

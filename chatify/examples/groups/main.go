package main

import (
	"encoding/json"
	"fmt"
	"gopher-playground/chatify/pkg/chatify"
	"gopher-playground/chatify/pkg/message"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

type CustomMessage struct {
	message.BaseMessage
}

func main() {
	server := chatify.NewServer(
		chatify.WithServerPath("/chat"),
	)

	group1 := server.NewGroup(
		chatify.WithGroupPath("/group1"),
		chatify.WithMessageFormat(parseMessage),
		chatify.WithGroupMiddleware(accessControl("fulano", "ciclano", "beltrano")),
	)

	group2 := server.NewGroup(
		chatify.WithGroupPath("/group2"),
		chatify.WithMessageFormat(parseMessage),
		chatify.WithGroupMiddleware(accessControl("mikaella", "geraldo")),
	)

	go server.Run()

	for {
		<-time.After(15 * time.Second)
		fmt.Println("Total chatters in group 1:", group1.TotalClients())
		fmt.Println("Total chatters in group 2:", group2.TotalClients())
	}
}

func accessControl(allowedUsers ...string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.GetHeader("user")
		fmt.Println(username)

		if !slices.Contains(allowedUsers, username) {
			ctx.String(http.StatusForbidden, "user not allowed in this group")
			ctx.Abort()
		}
	}
}

func parseMessage(data []byte) (message.Message, error) {
	// Marshal data
	var msg CustomMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return message.Message(&msg), nil
}

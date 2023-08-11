package chatify

import (
	"encoding/json"
	"fmt"
	"log"

	"gopher-playground/chatify/internal/connection"

	"github.com/gin-gonic/gin"
)

type ChatServer struct {
	clients   map[*Client]bool
	port      int
	path      string
	broadcast chan []byte
	format    func([]byte) (message, error)
}

func NewServer(options ...ChatServerOption) *ChatServer {
	s := &ChatServer{
		clients:   map[*Client]bool{},
		port:      8080,  //default
		path:      "/ws", // default
		broadcast: make(chan []byte),
		format: func(d []byte) (message, error) {
			var msg BaseMessage
			err := json.Unmarshal(d, &msg)
			return &msg, err
		},
	}

	// Add options to server
	for _, option := range options {
		option(s)
	}

	return s
}

func (s *ChatServer) Run() {
	r := gin.Default()
	upgrader := connection.NewWSUpgrader()

	r.GET(s.path, func(c *gin.Context) {
		// Upgrade http connection to websocket one
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatal("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		// Create and start new client
		NewClient(conn, s.broadcast, s.format).Start()
	})

	log.Printf("Chat server started at ws://localhost:%d/%s", s.port, s.path)
	r.Run(fmt.Sprintf(":%d", s.port))
}

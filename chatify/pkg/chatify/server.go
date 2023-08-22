package chatify

import (
	"fmt"
	"log"
	"sync"

	"gopher-playground/chatify/internal/connection"
	"gopher-playground/chatify/pkg/processor"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatServer struct {
	mutex     *sync.Mutex
	clients   map[*Client]bool
	port      int
	path      string
	broadcast chan []byte

	onConnect   func(conn *websocket.Conn)
	middlewares []gin.HandlerFunc

	// Message processor handlers
	format  *processor.Formatter
	persist *processor.Persistency
	customs []processor.Handler
}

func NewServer(options ...ChatServerOption) *ChatServer {
	s := &ChatServer{
		mutex:     &sync.Mutex{},
		clients:   map[*Client]bool{},
		port:      8080,  //default
		path:      "/ws", // default
		broadcast: make(chan []byte),
		onConnect: func(conn *websocket.Conn) {},
		format:    processor.NewDefaultFormatter(),
	}

	// Add options to server
	for _, option := range options {
		option(s)
	}

	return s
}

func (s *ChatServer) TotalClients() int {
	return len(s.clients)
}

func (s *ChatServer) Run() {
	r := gin.Default()
	upgrader := connection.NewWSUpgrader()

	// Create processor
	processor := processor.InitMsgProcessor(s.format, s.persist, s.customs...)

	handlers := append(s.middlewares, func(c *gin.Context) {
		// Upgrade http connection to websocket one
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatal("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		// On new connection callback
		s.onConnect(conn)

		// Create and start new client
		client := s.initClient(c, conn, processor)
		defer s.cleanClient(client)

		client.Start()
	})

	r.GET(s.path, handlers...)

	log.Printf("Chat server started at ws://localhost:%d/%s", s.port, s.path)
	r.Run(fmt.Sprintf(":%d", s.port))
}

func (s *ChatServer) initClient(c *gin.Context, conn *websocket.Conn, msgProcessor *processor.Processor) *Client {
	client := NewClient(c, conn, s.broadcast, msgProcessor)
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients[client] = true

	return client
}

func (s *ChatServer) cleanClient(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.clients, client)
}

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
	port        int
	path        string
	middlewares []gin.HandlerFunc
	router      *gin.Engine

	groups []*ChatGroup
}

type ChatGroup struct {
	mutex     *sync.Mutex
	clients   map[*Client]bool
	path      string
	broadcast chan []byte

	onConnect   func(conn *websocket.Conn)
	middlewares []gin.HandlerFunc
	router      *gin.RouterGroup

	// Message processor handlers
	format  *processor.Formatter
	persist *processor.Persistency
	customs []processor.Handler
}

func NewServer(options ...ChatServerOption) *ChatServer {
	s := &ChatServer{
		port:   8080,  //default
		path:   "/ws", // default
		router: gin.Default(),
	}

	// Add options to server
	for _, option := range options {
		option(s)
	}

	return s
}

func (s *ChatServer) NewGroup(options ...ChatGroupOption) *ChatGroup {
	g := &ChatGroup{
		mutex:     &sync.Mutex{},
		clients:   map[*Client]bool{},
		broadcast: make(chan []byte),
		onConnect: func(conn *websocket.Conn) {},
		format:    processor.NewDefaultFormatter(),
		path:      "/chat", // default
	}

	// Add options to server
	for _, option := range options {
		option(g)
	}

	// Setup router
	g.router = s.router.Group(s.path + g.path)

	// Append new group
	s.groups = append(s.groups, g)

	return g
}

func (s *ChatServer) Run() {
	// Set middleware
	s.router.Use(s.middlewares...)

	// Setup groups
	for _, g := range s.groups {
		g.setup()
	}

	log.Printf("Chat server started at ws://localhost:%d%s", s.port, s.path)
	s.router.Run(fmt.Sprintf(":%d", s.port))
}

func (s *ChatGroup) TotalClients() int {
	return len(s.clients)
}

func (s *ChatGroup) setup() {
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

	s.router.GET("", handlers...)
}

func (s *ChatGroup) initClient(c *gin.Context, conn *websocket.Conn, msgProcessor *processor.Processor) *Client {
	client := NewClient(c, conn, s.broadcast, msgProcessor)
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients[client] = true

	return client
}

func (s *ChatGroup) cleanClient(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.clients, client)
}

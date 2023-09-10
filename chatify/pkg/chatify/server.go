package chatify

import (
	"context"
	"fmt"
	"log"
	"sync"

	"gopher-playground/chatify/internal/async"
	"gopher-playground/chatify/internal/connection"
	"gopher-playground/chatify/pkg/processor"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	broadcast async.Broadcaster

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
		broadcast: *async.NewBroadcaster(make(chan []byte)),
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
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, g := range s.groups {
		wg.Add(1)
		g.setup(&wg, ctx, cancel)
	}

	log.Printf("Chat server started at ws://localhost:%d%s", s.port, s.path)
	s.router.Run(fmt.Sprintf(":%d", s.port))

	// End groups
	cancel()
	wg.Wait()
}

func (s *ChatGroup) TotalClients() int {
	return len(s.clients)
}

func (s *ChatGroup) setup(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
	upgrader := connection.NewWSUpgrader()

	// Create processor
	processor := processor.InitMsgProcessor(s.format, s.persist, s.customs...)

	// Init broadcaster
	go async.RunWorker(s.broadcast.Receiver, wg, ctx, cancel, s.broadcast.Start)

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
	id := uuid.New().String()
	dataChannel := s.broadcast.Register(id)

	client := NewClient(id, dataChannel, c, conn, s.broadcast, msgProcessor)
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients[client] = true

	return client
}

func (s *ChatGroup) cleanClient(client *Client) {
	s.broadcast.Unregister(client.Id)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.clients, client)
}

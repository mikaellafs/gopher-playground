package chatify

import (
	"gopher-playground/chatify/pkg/message"
	"gopher-playground/chatify/pkg/processor"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatServerOption func(server *ChatServer)
type ChatGroupOption func(server *ChatGroup)

// Option to set custom format function
func WithMessageFormat(format func([]byte) (message.Message, error)) ChatGroupOption {
	return func(server *ChatGroup) {
		server.format = processor.NewMessageFormatter(format)
	}
}

// Option to set custom message persistence
func WithMessageStore(store message.Store) ChatGroupOption {
	return func(server *ChatGroup) {
		server.persist = processor.NewMessagePersister(store)
	}
}

func WithCustomMessageHandler(handler processor.Handler) ChatGroupOption {
	return func(server *ChatGroup) {
		server.customs = append(server.customs, handler)
	}
}

// Option to set custom port
func WithPort(port int) ChatServerOption {
	return func(server *ChatServer) {
		server.port = port
	}
}

func WithServerPath(path string) ChatServerOption {
	return func(server *ChatServer) {
		server.path = path
	}
}

func WithGroupPath(path string) ChatGroupOption {
	return func(server *ChatGroup) {
		server.path = path
	}
}

func WithOnConnectionCallback(onConnect func(*websocket.Conn)) ChatGroupOption {
	return func(server *ChatGroup) {
		server.onConnect = onConnect
	}
}

func WithServerMiddleware(mid gin.HandlerFunc) ChatServerOption {
	return func(server *ChatServer) {
		server.middlewares = append(server.middlewares, mid)
	}
}

func WithGroupMiddleware(mid gin.HandlerFunc) ChatGroupOption {
	return func(server *ChatGroup) {
		server.middlewares = append(server.middlewares, mid)
	}
}

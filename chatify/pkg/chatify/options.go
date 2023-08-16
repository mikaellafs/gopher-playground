package chatify

import (
	"gopher-playground/chatify/pkg/message"
	"gopher-playground/chatify/pkg/processor"
)

type ChatServerOption func(server *ChatServer)

// Option to set custom format function
func WithMessageFormat(format func([]byte) (message.Message, error)) ChatServerOption {
	return func(server *ChatServer) {
		server.format = processor.NewMessageFormatter(format)
	}
}

// Option to set custom message persistence
func WithMessageStore(store message.Store) ChatServerOption {
	return func(server *ChatServer) {
		server.persist = processor.NewMessagePersister(store)
	}
}

func WithCustomMessageHandler(handler processor.Handler) ChatServerOption {
	return func(server *ChatServer) {
		server.customs = append(server.customs, handler)
	}
}

// Option to set custom port
func WithPort(port int) ChatServerOption {
	return func(server *ChatServer) {
		server.port = port
	}
}

func WithPath(path string) ChatServerOption {
	return func(server *ChatServer) {
		server.path = path
	}
}

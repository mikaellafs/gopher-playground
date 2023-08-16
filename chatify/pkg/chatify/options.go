package chatify

type ChatServerOption func(server *ChatServer)

// Option to set custom format function
func WithMessageFormat(format func([]byte) (message, error)) ChatServerOption {
	return func(server *ChatServer) {
		server.format = format
	}
}

// Option to set custom message persistence
func WithMessageStore(store MessageStore) ChatServerOption {
	return func(server *ChatServer) {
		server.messageStore = store
	}
}

// Option to set custom port
func WithPort(port int) ChatServerOption {
	return func(server *ChatServer) {
		server.port = port
	}
}

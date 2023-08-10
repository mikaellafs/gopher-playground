package chatify

type ChatServerOption func(server *ChatServer)

// Option to set custom callback function
// func WithCallback(callback func(Message)) ChatServerOption {
// 	return func(server *ChatServer) {
// 		server.callback = callback
// 	}
// }

// // Option to set custom message persistence
// func WithMessageStore(store MessageStore) ChatServerOption {
// 	return func(server *ChatServer) {
// 		server.messageStore = store
// 	}
// }

// Option to set custom port
func WithPort(port int) ChatServerOption {
	return func(server *ChatServer) {
		server.port = port
	}
}

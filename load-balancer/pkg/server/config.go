package server

import "time"

type ConfigOption func(*ServerPool)

func WithTimeout(timeout time.Duration) ConfigOption {
	return func(sp *ServerPool) {
		sp.timeout = timeout
	}
}

// Max connections servers can handle before creating a new instance, default is 200
func WithMaxConnections(max int) ConfigOption {
	return func(sp *ServerPool) {
		sp.maxConnections = max
	}
}

// Min connections servers have to handle before killing instance, default is 40
func WithMinConnections(min int) ConfigOption {
	return func(sp *ServerPool) {
		sp.minConnections = min
	}
}

// Minimum amount of servers to keep running, default is 1
func WithMinServers(n int) ConfigOption {
	return func(sp *ServerPool) {
		sp.minServers = n
	}
}

// Max attempts to health check before killing an instance, default is 3
func WithMaxRetries(n int) ConfigOption {
	return func(sp *ServerPool) {
		sp.maxRetries = n
	}
}

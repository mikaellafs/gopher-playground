package server

type ServerPool struct {
	service Service
}

func NewServerPool(service Service, configs ...ConfigOption) *ServerPool {
	s := &ServerPool{
		service: service,
	}

	// Apply config options
	for _, apply := range configs {
		apply(s)
	}

	return s
}

func (p *ServerPool) Shutdown() error {
	return nil
}

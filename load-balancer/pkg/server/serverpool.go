package server

import (
	"gopher-playground/load-balancer/internal/utils"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type ServerPool struct {
	mutex   *sync.Mutex
	servers map[string]*Server
	service Service

	timeout        time.Duration
	maxConnections int
	minConnections int
	minServers     int
	maxRetries     int
}

func NewServerPool(service Service, configs ...ConfigOption) *ServerPool {
	s := &ServerPool{
		mutex:   &sync.Mutex{},
		servers: map[string]*Server{},
		service: service,

		timeout:        10 * time.Second,
		maxConnections: 200,
		minConnections: 40,
		minServers:     1,
		maxRetries:     3,
	}

	// Apply config options
	for _, apply := range configs {
		apply(s)
	}

	return s
}

func (p *ServerPool) Shutdown() error {
	var wg sync.WaitGroup
	wg.Add(len(p.servers))

	errMutex := &sync.Mutex{}
	var errs []error

	for id, s := range p.servers {
		go p.shutdown(&wg, id, s, &errs, errMutex)
	}

	wg.Wait()

	return utils.MergeErrors(errs)
}

func (p *ServerPool) shutdown(wg *sync.WaitGroup, id string, s *Server, errs *[]error, errMutex *sync.Mutex) {
	defer wg.Done()

	var err error
	for i := 0; i < 3; i++ {
		if err = p.service.Shutdown(s); err == nil {
			return
		}

		log.Printf("Failed to shutdown server %s. Attempting again in %d seconds\n", id, 5)
		time.Sleep(5 * time.Second)
	}

	errMutex.Lock()
	defer errMutex.Unlock()

	*errs = append(*errs, errors.Wrap(err, "failed to shutdown server "+id))
}

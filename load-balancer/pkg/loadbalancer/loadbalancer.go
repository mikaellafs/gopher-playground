package loadbalancer

import (
	"gopher-playground/load-balancer/pkg/server"
	"sync"
)

type LoadBalancer struct {
	mutex *sync.Mutex
	pools map[string]*server.ServerPool
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		mutex: &sync.Mutex{},
		pools: map[string]*server.ServerPool{},
	}
}

func (l *LoadBalancer) RegisterService(name string, service server.Service, configs ...server.ConfigOption) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pools[name] != nil {
		return ErrServiceNameAlreadyExists
	}

	l.pools[name] = server.NewServerPool(service, configs...)
	return nil
}

func (l *LoadBalancer) UnregisterService(name string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.pools[name] == nil {
		return ErrServiceNotExists
	}

	err := l.pools[name].Shutdown()
	if err != nil {
		return err
	}

	l.pools[name] = nil
	return nil
}

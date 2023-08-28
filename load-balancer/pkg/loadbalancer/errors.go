package loadbalancer

import "errors"

var (
	ErrServiceNameAlreadyExists = errors.New("service name already used")
	ErrServiceNotExists         = errors.New("service does not exist")
)

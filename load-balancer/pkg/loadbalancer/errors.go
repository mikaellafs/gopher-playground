package loadbalancer

import "errors"

var (
	ErrServiceNameAlreadyExists = errors.New("service name already used")
)

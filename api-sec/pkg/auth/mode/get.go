package mode

import (
	"errors"
	"strings"
)

const (
	basic string = "basic"
)

var (
	ErrInvalidAuthMode = errors.New("Invalid auth mode.")
)

func Get(mode string) (AuthMode, error) {
	mode = strings.ToLower(mode)

	switch mode {
	case basic:
		return &BasicAuth{}, nil
	}

	return nil, ErrInvalidAuthMode
}

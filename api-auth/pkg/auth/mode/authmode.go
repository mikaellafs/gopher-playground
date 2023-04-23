package mode

import (
	"errors"
	"gopher-playground/api-auth/pkg/auth/user"
)

var (
	ErrInvalidAuthHeader  = errors.New("Invalid auth header")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type AuthMode interface {
	Authenticate(authHeader string, userRepo user.Repository) (*user.User, error)
}

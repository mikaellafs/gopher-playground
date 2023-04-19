package auth

import (
	"errors"
	"gopher-playground/api-auth/pkg/user"
)

var (
	ErrInvalidAuthHeader  = errors.New("Invalid auth header")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type AuthMode interface {
	Authenticate(authHeader string, userRepo user.Repository) (string, error)
}

package mode

import (
	"errors"

	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidAuthHeader  = errors.New("Invalid auth header")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

type AuthMode interface {
	Authenticate(authHeader string, userRepo user.Repository) (*user.User, error)
	GenerateToken(tstore token.TokenStore, c *gin.Context) *token.Token
}

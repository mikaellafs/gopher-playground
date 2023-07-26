package mode

import (
	"errors"

	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

var (
	ErrMissingAuthHeader  = errors.New("missing auth header")
	ErrInvalidAuthHeader  = errors.New("invalid auth header")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthMode interface {
	Authenticate(c *gin.Context) (*user.User, error)
	GenerateToken(c *gin.Context, username string) (*token.Token, error)
	Logout(c *gin.Context)
}

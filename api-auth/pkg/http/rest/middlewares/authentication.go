package middlewares

import (
	"gopher-playground/api-auth/pkg/auth"
	"gopher-playground/api-auth/pkg/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Authentication struct {
	repo user.Repository
	mode auth.AuthMode
}

func NewAuthenticator(authMode auth.AuthMode) *Authentication {
	return &Authentication{
		mode: authMode,
	}
}

func (m *Authentication) Middleware(c *gin.Context) {
	authH := c.GetHeader("Authorization")
	if authH == "" {
		c.Next()
	}

	username, err := m.mode.Authenticate(authH, m.repo)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	c.Set("user", username)
	c.Next()

}

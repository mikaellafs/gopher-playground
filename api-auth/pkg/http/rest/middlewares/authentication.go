package middlewares

import (
	authmode "gopher-playground/api-auth/pkg/auth/mode"
	"gopher-playground/api-auth/pkg/auth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Authentication struct {
	repo user.Repository
	mode authmode.AuthMode
}

func NewAuthenticator(authMode authmode.AuthMode, repo user.Repository) *Authentication {
	return &Authentication{
		mode: authMode,
		repo: repo,
	}
}

func (m *Authentication) Middleware(c *gin.Context) {
	authH := c.GetHeader("Authorization")
	if authH == "" {
		c.Next()
	}

	user, err := m.mode.Authenticate(authH, m.repo)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	c.Set("user", user)
	c.Next()
}

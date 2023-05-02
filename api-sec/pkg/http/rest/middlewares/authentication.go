package middlewares

import (
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	"gopher-playground/api-sec/pkg/auth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(authMode authmode.AuthMode, repo user.Repository) func(*gin.Context) {
	return func(c *gin.Context) {
		authH := c.GetHeader("Authorization")
		if authH == "" {
			c.Next()
			return
		}

		user, err := authMode.Authenticate(authH, repo)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
		}

		c.Set("user", user)
		c.Next()
	}
}

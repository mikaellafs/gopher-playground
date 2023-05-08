package middlewares

import (
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	"gopher-playground/api-sec/pkg/auth/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

const authErrKeyName string = "auth-error"

func Authentication(authMode authmode.AuthMode, repo user.Repository) func(*gin.Context) {
	return func(c *gin.Context) {
		authH := c.GetHeader("Authorization")
		if authH == "" {
			c.Next()
			return
		}

		user, err := authMode.Authenticate(authH, repo)
		if err != nil {
			c.Set(authErrKeyName, err.Error())
			c.Next()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func RequireAuth() func(*gin.Context) {
	return func(c *gin.Context) {
		errMsg, exists := c.Get(authErrKeyName)
		if !exists {
			c.Next()
			return
		}

		msg, _ := errMsg.(string)
		c.String(http.StatusUnauthorized, msg)
		c.Abort()
	}
}

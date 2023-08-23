package middlewares

import (
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	"net/http"

	"github.com/gin-gonic/gin"
)

const authErrKeyName string = "auth-error"

func Authentication(authMode authmode.AuthMode) func(*gin.Context) {
	return func(c *gin.Context) {
		user, err := authMode.Authenticate(c)
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

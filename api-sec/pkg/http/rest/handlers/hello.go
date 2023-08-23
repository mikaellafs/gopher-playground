package handlers

import (
	"net/http"

	"gopher-playground/api-sec/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

func Hello() func(*gin.Context) {
	return func(c *gin.Context) {
		u, exists := c.Get("user")
		if !exists {
			c.String(http.StatusBadRequest, "Missing user")
			c.Abort()
		}

		user, _ := u.(*user.User)

		c.String(http.StatusOK, "Hi, "+user.Name+"!")
	}
}

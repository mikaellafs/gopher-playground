package handlers

import (
	"net/http"

	"gopher-playground/api-auth/pkg/auth/user"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Say(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.String(http.StatusBadRequest, "Missing user")
		c.Abort()
	}

	user, _ := u.(*user.User)

	c.String(http.StatusOK, "Hi, "+user.Name+"!")
}

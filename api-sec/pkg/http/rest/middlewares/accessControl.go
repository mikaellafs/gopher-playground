package middlewares

import (
	"gopher-playground/api-sec/pkg/auth/accesscontrol"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccessControl(ac accesscontrol.AccessControl) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get user, domain, resource and action
		user := getUsernameFromContext(c)
		// domain := c.GetString("domain")
		path := c.FullPath()
		method := c.Request.Method

		// Enforce
		if allowed, err := ac.Enforce(user, path, method); err == nil && !allowed {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		} else if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

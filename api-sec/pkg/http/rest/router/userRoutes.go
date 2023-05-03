package router

import (
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setUserRoutes(rg *gin.RouterGroup, r user.Repository) {
	userGroup := rg.Group("/users")

	userGroup.POST("", handlers.CreateUser(r))

	// userGroup.DELETE("", handlers.DeleteUser(r))
}

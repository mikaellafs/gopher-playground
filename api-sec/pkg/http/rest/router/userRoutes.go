package router

import (
	ac "gopher-playground/api-sec/pkg/auth/accesscontrol"
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setUserRoutes(rg *gin.RouterGroup, r user.Repository, ac ac.AccessControl) {
	userGroup := rg.Group("/users")

	userGroup.POST("", handlers.CreateUser(r))

	// userGroup.DELETE("", handlers.DeleteUser(r))
}

package router

import (
	"gopher-playground/api-auth/pkg/auth/user"
	"gopher-playground/api-auth/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setUserRoutes(rg *gin.RouterGroup, r user.Repository) {
	userGroup := rg.Group("/users")

	userHandler := handlers.NewUserHandler(r)

	userGroup.POST("", userHandler.CriarUsuario)
}

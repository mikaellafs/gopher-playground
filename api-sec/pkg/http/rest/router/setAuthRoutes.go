package router

import (
	"gopher-playground/api-sec/pkg/auth/mode"
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setAuthRoutes(rg *gin.RouterGroup, repo user.Repository, mode mode.AuthMode) {
	rg.POST("/login", handlers.Login(repo, mode))
	rg.POST("/logout", handlers.Logout(mode))

	rg.POST("/token/refresh", handlers.Refresh(mode))
}

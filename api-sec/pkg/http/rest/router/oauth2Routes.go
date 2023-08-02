package router

import (
	"gopher-playground/api-sec/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setOauth2Routes(rg *gin.RouterGroup) {
	rg.GET("/google/login", handlers.HandleGoogleLogin)
	rg.GET("/google/callback", handlers.HandleGoogleCallback)
}

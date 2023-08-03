package router

import (
	"gopher-playground/api-sec/pkg/auth/accesscontrol"
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setOauth2Routes(rg *gin.RouterGroup, repo user.Repository, authMode authmode.AuthMode, ac accesscontrol.AccessControl) {
	rg.GET("/google/login", handlers.HandleGoogleLogin)
	rg.GET("/google/callback", handlers.HandleGoogleCallback(repo, authMode, ac))
}

package router

import (
	ac "gopher-playground/api-sec/pkg/auth/accesscontrol"
	"gopher-playground/api-sec/pkg/http/rest/handlers"
	"gopher-playground/api-sec/pkg/http/rest/middlewares"
	"gopher-playground/api-sec/pkg/log"

	"github.com/gin-gonic/gin"
)

func setLogsRoute(rg *gin.RouterGroup, r log.Repository, ac ac.AccessControl) {
	logsGroup := rg.Group("/logs")

	// Set acces control to route
	logsGroup.Use(middlewares.AccessControl(ac))

	logsGroup.GET("", handlers.ListAllLogs(r))
}

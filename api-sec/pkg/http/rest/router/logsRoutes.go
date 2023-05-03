package router

import (
	"gopher-playground/api-sec/pkg/http/rest/handlers"
	"gopher-playground/api-sec/pkg/log"

	"github.com/gin-gonic/gin"
)

func setLogsRoute(rg *gin.RouterGroup, r log.Repository) {
	logsGroup := rg.Group("/logs")

	logsGroup.GET("", handlers.ListAllLogs(r))
}

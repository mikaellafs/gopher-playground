package router

import (
	"gopher-playground/api-auth/pkg/http/rest/handlers"
	"gopher-playground/api-auth/pkg/log"

	"github.com/gin-gonic/gin"
)

func setLogsRoute(rg *gin.RouterGroup, r log.Repository) {
	logsGroup := rg.Group("/logs")

	lhandler := handlers.NewLogsHandler(r)
	logsGroup.GET("", lhandler.ListAll)
}

package router

import (
	"gopher-playground/api-auth/pkg/http/rest/handlers"

	"github.com/gin-gonic/gin"
)

func setHelloRoutes(rg *gin.RouterGroup) {
	helloGroup := rg.Group("/hello")

	helloHandler := handlers.NewHelloHandler()

	helloGroup.GET("", helloHandler.Say)
}

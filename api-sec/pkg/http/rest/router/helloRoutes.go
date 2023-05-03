package router

import (
	ac "gopher-playground/api-sec/pkg/auth/accesscontrol"
	"gopher-playground/api-sec/pkg/http/rest/handlers"
	"gopher-playground/api-sec/pkg/http/rest/middlewares"

	"github.com/gin-gonic/gin"
)

func setHelloRoutes(rg *gin.RouterGroup, ac ac.AccessControl) {
	helloGroup := rg.Group("/hello")

	// Set acces control to route
	helloGroup.Use(middlewares.AccessControl(ac))

	helloGroup.GET("", handlers.Hello())
}

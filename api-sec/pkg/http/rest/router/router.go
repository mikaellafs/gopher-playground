package router

import (
	"net/http"
	"time"

	ac "gopher-playground/api-sec/pkg/auth/accesscontrol"
	authmode "gopher-playground/api-sec/pkg/auth/mode"
	"gopher-playground/api-sec/pkg/auth/token"
	"gopher-playground/api-sec/pkg/auth/user"
	"gopher-playground/api-sec/pkg/http/rest/middlewares"
	"gopher-playground/api-sec/pkg/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type middleware func(c *gin.Context)

type Config struct {
	RateLimit  int
	RetryAfter float64

	AuthMode      authmode.AuthMode
	AccessControl ac.AccessControl
	TokenStore    token.TokenStore

	UserRepo user.Repository
	LogRepo  log.Repository
}

func Initialize(cfg *Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
			// "http://localhost:3000",      // Dev address
			// "http://localhost",           // Test addresses
			// os.Getenv("FRONTEND_ORIGIN"), // Address of where the front-end is deployed
		},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	rg := r.Group("/api")

	rg.GET("/", healthCheck)

	setGlobalMiddlewares(rg, cfg)

	setAuthRoutes(rg, cfg.UserRepo, cfg.AuthMode)
	setUserRoutes(rg, cfg.UserRepo, cfg.AccessControl)
	setHelloRoutes(rg, cfg.AccessControl)
	setLogsRoute(rg, cfg.LogRepo, cfg.AccessControl)

	return r
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func setGlobalMiddlewares(rg *gin.RouterGroup, cfg *Config) {
	rg.Use(middlewares.RateLimiting(cfg.RateLimit, cfg.RetryAfter))
	rg.Use(middlewares.Authentication(cfg.AuthMode))
	rg.Use(middlewares.StartAuditLog(cfg.LogRepo))

	setPostMiddleware(rg, middlewares.EndAuditLog(cfg.LogRepo))
}

// Middleware that need to be execute after handler
func setPostMiddleware(rg *gin.RouterGroup, m middleware) {
	rg.Use(func(c *gin.Context) {
		c.Next()

		m(c)
	})
}

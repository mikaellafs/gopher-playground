package router

import (
	"net/http"
	"time"

	authmode "gopher-playground/api-sec/pkg/auth/mode"
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

	AuthMode authmode.AuthMode

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

	setMiddlewares(rg, cfg)

	setUserRoutes(rg, cfg.UserRepo)
	setHelloRoutes(rg)
	setLogsRoute(rg, cfg.LogRepo)

	return r
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func setMiddlewares(rg *gin.RouterGroup, cfg *Config) {
	rg.Use(middlewares.NewRateLimiting(cfg.RateLimit, cfg.RetryAfter).Middleware)
	rg.Use(middlewares.NewAuthenticator(cfg.AuthMode, cfg.UserRepo).Middleware)

	logMiddleware := middlewares.NewAuditLog(cfg.LogRepo)
	rg.Use(logMiddleware.StartMiddleware)

	setPostMiddleware(rg, logMiddleware.EndMiddleware)
}

// Middleware that need to be execute after handler
func setPostMiddleware(rg *gin.RouterGroup, m middleware) {
	rg.Use(func(c *gin.Context) {
		c.Next()

		m(c)
	})
}

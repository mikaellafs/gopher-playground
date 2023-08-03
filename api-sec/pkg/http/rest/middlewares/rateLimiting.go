package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiting(rateLimit int, retryAfter float64) func(*gin.Context) {
	limiter := rate.NewLimiter(rate.Every(time.Second), rateLimit)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.Writer.Header().Set("Retry-After", strconv.FormatFloat(retryAfter, 'E', -1, 64))
			c.Status(http.StatusTooManyRequests) // 429
		}

		c.Next()
	}
}

package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type RateLimiting struct {
	limiter    *rate.Limiter
	retryAfter float64
}

func NewRateLimiting(rateLimit int, retryAfter float64) *RateLimiting {
	return &RateLimiting{
		retryAfter: retryAfter,
		limiter:    rate.NewLimiter(rate.Every(time.Second), rateLimit),
	}
}
func (r *RateLimiting) Middleware(c *gin.Context) {
	if !r.limiter.Allow() {
		c.Writer.Header().Set("Retry-After", strconv.FormatFloat(r.retryAfter, 'E', -1, 64))
		c.Status(http.StatusTooManyRequests) // 429
	}

	c.Next()
}

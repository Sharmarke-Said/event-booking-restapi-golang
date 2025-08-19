package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// GinRateLimiter returns a Gin middleware that rate-limits requests based on a token bucket algorithm.
// It uses a global limiter to apply the same rate limit to all requests.
func GinRateLimiter() gin.HandlerFunc {
	// A token bucket with a capacity of 5 and a replenishment rate of 1 token per second.
	// This means it allows bursts of up to 5 requests, and then 1 request per second.
	limiter := rate.NewLimiter(rate.Every(time.Second), 5)
	
	return func(context *gin.Context) {
		if !limiter.Allow() {
			log.Println("Rate limit exceeded for request from", context.ClientIP())
			context.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
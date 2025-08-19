package router

import (
	"example.com/event-booking-restapi/internal/event"
	"example.com/event-booking-restapi/internal/middleware"
	"example.com/event-booking-restapi/internal/user"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API routes with versioning and grouping.
func RegisterRoutes(server *gin.Engine) {
    // Create an API version group, e.g., /api/v1
    v1 := server.Group("/api/v1")

    // Apply rate limiting middleware to all v1 endpoints
    v1.Use(middleware.GinRateLimiter())

    // Delegate route registration for each feature to its own package,
    // passing the versioned router group.
    user.RegisterRoutes(v1)
    event.RegisterRoutes(v1)
}
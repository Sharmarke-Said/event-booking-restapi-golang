package event

import (
	"example.com/event-booking-restapi/internal/auth/roles"
	"example.com/event-booking-restapi/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all event-related API routes on the provided router group.
func RegisterRoutes(router *gin.RouterGroup) {
    // Public route
    router.GET("/events", getEvents)

    // Group authenticated routes and apply the middleware
    authenticated := router.Group("/events")
    authenticated.Use(middleware.Protect())

    authenticated.POST("", createEvent)
    authenticated.GET("/:id", getEvent)
    authenticated.PUT("/:id", updateEvent)
    authenticated.DELETE("/:id", deleteEvent)
    authenticated.POST("/:id/register", registerForEvent)
    authenticated.DELETE("/:id/register", cancelRegistration)

    adminRegistrationsRouter := router.Group("/events/:id/registrations")
    adminRegistrationsRouter.Use(middleware.Protect(), middleware.RestrictTo(roles.Admin))
    {
        adminRegistrationsRouter.GET("", GetEventRegistrations)
    }
    
}
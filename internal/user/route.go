package user

import (
	"example.com/event-booking-restapi/internal/auth/roles"
	"example.com/event-booking-restapi/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all user-related API routes on the provided router group.
func RegisterRoutes(router *gin.RouterGroup) {
    // Public routes
    router.POST("/signup", signup)
    router.POST("/login", login)

    // Admin-only routes for user management
    adminRouter := router.Group("/users")
    adminRouter.Use(middleware.Protect(), middleware.RestrictTo(roles.Admin))
    {
        adminRouter.GET("", GetUsers)
        adminRouter.GET("/:id", GetUser)
        adminRouter.PATCH("/:id", UpdateUser)
        adminRouter.DELETE("/:id", DeleteUser)
    }
}
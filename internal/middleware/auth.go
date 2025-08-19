package middleware

import (
	"net/http"

	"example.com/event-booking-restapi/internal/auth"
	"github.com/gin-gonic/gin"
)

func Protect() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHeader := context.Request.Header.Get("Authorization")
		   if tokenHeader == "" {
			   context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
				 context.Abort()
			   return
		   }
		token := tokenHeader[len("Bearer "):]
		   if token == "" {
			   context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
				 context.Abort()
			   return
		   }

		   claims, err := auth.VerifyToken(token)
		   if err != nil {
			   context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
				 context.Abort()
			   return
		   }

		// Extract userId and role from claims and set it in the context
		userId := claims["userId"].(float64)
		role := claims["role"].(string)

		context.Set("userId", int64(userId))
		context.Set("role", role)

		context.Next()
	}
}

// RestrictToRole restricts access to certain roles
func RestrictTo(role string) gin.HandlerFunc {
	return  func(context *gin.Context) {
		userRole, exists := context.Get("role")
		   if !exists || userRole != role {
			   context.JSON(http.StatusForbidden, gin.H{"message": "You're not authorized to perform this action"})
				 context.Abort()
			   return
		   }
		context.Next()
	}
}

// func RestrictToRole(allowedRoles ...string) gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		role, exists := context.Get("role")
// 		if !exists {
// 			errors.HandleAPIError(context, errors.NewAPIError(http.StatusForbidden, "Forbidden"))
// 			return
// 		}

// 		for _, allowedRole := range allowedRoles {
// 			if role == allowedRole {
// 				context.Next()
// 				return
// 			}
// 		}

// 		errors.HandleAPIError(context, errors.NewAPIError(http.StatusForbidden, "Forbidden"))
// 	}
// 	}
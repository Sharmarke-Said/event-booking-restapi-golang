package errors

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // APIError represents a custom error with a status code and message.
// type APIError struct {
//     StatusCode int
//     Message    string
// }

// func (e APIError) Error() string {
//     return e.Message
// }

// // NewAPIError creates a new APIError.
// func NewAPIError(statusCode int, message string) *APIError {
//     return &APIError{
//         StatusCode: statusCode,
//         Message:    message,
//     }
// }

// // HandleAPIError is a centralized function to handle and respond with API errors.
// func HandleAPIError(context *gin.Context, err error) {
//     if apiErr, ok := err.(*APIError); ok {
//         context.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
//         return
//     }
//     // For unhandled internal errors
//     context.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred."})
// }
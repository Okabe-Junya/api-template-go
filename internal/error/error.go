package apierror

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents the error response format for API
type ErrorResponse struct {
	Message string `json:"message"` // 8-byte aligned string first
	Code    int    `json:"code"`    // 4-byte int next
}

// Error implements the error interface
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("status %d: %s", e.Code, e.Message)
}

// API error definitions
var (
	ErrBadRequest          = &ErrorResponse{Code: http.StatusBadRequest, Message: "Bad request"}
	ErrUnauthorized        = &ErrorResponse{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden           = &ErrorResponse{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound            = &ErrorResponse{Code: http.StatusNotFound, Message: "Not found"}
	ErrMethodNotAllowed    = &ErrorResponse{Code: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrConflict            = &ErrorResponse{Code: http.StatusConflict, Message: "Conflict"}
	ErrInternalServerError = &ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error"}
)

// NewErrorWithMessage creates an error response with the specified status code and message
func NewErrorWithMessage(statusCode int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    statusCode,
		Message: message,
	}
}

// RespondWithError sends an error response in JSON format
func RespondWithError(c *gin.Context, err *ErrorResponse) {
	c.JSON(err.Code, err)
}

// ErrorHandlingMiddleware catches errors and returns appropriate responses
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Return appropriate response based on error type
			if apiErr, ok := err.Err.(*ErrorResponse); ok {
				RespondWithError(c, apiErr)
				return
			}

			// Default to server error
			RespondWithError(c, ErrInternalServerError)
			return
		}
	}
}

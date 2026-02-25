package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Handle panics
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse[any]("Internal server error"))
			c.Abort()
		}
	}
}

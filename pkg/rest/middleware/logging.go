package middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Info().Interface("request", c.Request.Body).Msg("request")
		c.Next()
		// log.Info().Interface("response", c.Writer.Status()).Msg("response")
	}
}

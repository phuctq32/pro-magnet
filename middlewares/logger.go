package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// handler
		c.Next()

		latency := time.Since(startTime)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}

		raw := c.Request.URL.RawQuery
		path := c.Request.URL.Path
		if raw != "" {
			path = path + "?" + raw
		}

		log.Info().
			Int("statusCode", c.Writer.Status()).
			Str("clientIP", c.ClientIP()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("latency", latency.String()).
			Msg("http request")
	}
}

package vibes

import (
	"time"

	"github.com/gin-gonic/gin"
)

// EmotionalLoggingMiddleware returns a middleware that logs requests with emotional vibes
func EmotionalLoggingMiddleware(logger *VibeLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// start timer
		startTime := time.Now()

		// process request
		c.Next()

		// log after request
		latency := time.Since(startTime)
		status := c.Writer.Status()
		path := c.Request.URL.Path
		method := c.Request.Method

		// log differently based on status
		switch {
		case status >= 500:
			logger.Crap("%s %s [%d] (%s) %s", method, path, status, latency, GetStatusEmoji(status))
		case status >= 400:
			logger.Uhoh("%s %s [%d] (%s) %s", method, path, status, latency, GetStatusEmoji(status))
		case status >= 300:
			logger.Uhoh("%s %s [%d] (%s) %s", method, path, status, latency, GetStatusEmoji(status))
		default:
			logger.Fyi("%s %s [%d] (%s) %s", method, path, status, latency, GetStatusEmoji(status))
		}
	}
}

package vibes

import (
	"github.com/gin-gonic/gin"
)

const (
	// XStatusEmoji is the header key for the emoji status representation
	XStatusEmoji = "X-Status-Emoji"
)

// EmojiStatusMiddleware returns a middleware that injects the emoji status into responses
func EmojiStatusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		vibesWriter := NewVibesResponseWriter(c)
		c.Writer = vibesWriter

		c.Next()
	}
}

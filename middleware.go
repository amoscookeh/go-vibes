package vibes

import (
	"github.com/gin-gonic/gin"
)

const (
	// XStatusEmoji is the header key for the emoji status representation
	XStatusEmoji = "X-Status-Emoji"
)

// Middleware type defines a middleware handler
type Middleware func(c *Context)

// wrapMiddleware wraps a Vibes middleware to a Gin middleware
func wrapMiddleware(m Middleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		m(&Context{Context: c})
	}
}

// VibeyMethodConverterMiddleware converts standard HTTP methods to vibey methods
// This is useful for HTTP clients that still use standard methods
func VibeyMethodConverterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the original method
		method := c.Request.Method

		// Convert to vibey method
		vibeyMethod := FromHTTPMethod(method)

		// Set a header to indicate the conversion
		c.Header("X-Original-Method", method)
		c.Header("X-Vibey-Method", string(vibeyMethod))

		// Continue with the request
		c.Next()
	}
}

// EmojiStatusMiddleware returns a middleware that injects the emoji status into responses
func EmojiStatusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		vibesWriter := NewVibesResponseWriter(c)
		c.Writer = vibesWriter

		c.Next()
	}
}

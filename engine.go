package vibes

import (
	"github.com/gin-gonic/gin"
)

// VibesEngine wraps the Gin engine to provide the vibes functionality
type VibesEngine struct {
	*gin.Engine
}

// New creates a new VibesEngine with default settings
func New() *VibesEngine {
	engine := gin.New()

	vibesEngine := &VibesEngine{
		Engine: engine,
	}

	// Apply the emoji status middleware by default
	engine.Use(EmojiStatusMiddleware())

	// Add default middleware from gin
	engine.Use(gin.Recovery())

	return vibesEngine
}

// Default creates a new VibesEngine with the default Gin middleware
func Default() *VibesEngine {
	engine := gin.Default()

	vibesEngine := &VibesEngine{
		Engine: engine,
	}

	// inject emoji middleware at the beginning
	middlewares := engine.Handlers
	engine.Handlers = make([]gin.HandlerFunc, 0, len(middlewares)+1)
	engine.Use(EmojiStatusMiddleware())
	for _, middleware := range middlewares {
		if middleware != nil {
			engine.Use(middleware)
		}
	}

	return vibesEngine
}

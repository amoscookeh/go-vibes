package vibes

import (
	"github.com/gin-gonic/gin"
)

func init() {
	SetReleaseMode()
}

type VibesEngine struct {
	*gin.Engine
	Logger *VibeLogger
}

// disables Gin's debug logs
func SetReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

func New() *VibesEngine {
	engine := gin.New()

	vibesEngine := &VibesEngine{
		Engine: engine,
		Logger: DefaultVibeLogger(),
	}

	// apply vibes middleware
	engine.Use(EmojiStatusMiddleware())
	engine.Use(EmotionalLoggingMiddleware(vibesEngine.Logger))
	engine.Use(gin.Recovery())

	vibesEngine.Logger.Fyi("go-vibes engine created with extra vibes ✨")

	return vibesEngine
}

// Default creates a new VibesEngine with the default Gin middleware
func Default() *VibesEngine {
	engine := gin.Default()

	vibesEngine := &VibesEngine{
		Engine: engine,
		Logger: DefaultVibeLogger(),
	}

	// inject emoji middleware at the beginning
	middlewares := engine.Handlers
	engine.Handlers = make([]gin.HandlerFunc, 0, len(middlewares)+1)

	// add vibes middleware first
	engine.Use(EmojiStatusMiddleware())
	engine.Use(EmotionalLoggingMiddleware(vibesEngine.Logger))

	// add other middlewares back
	for _, middleware := range middlewares {
		if middleware != nil {
			engine.Use(middleware)
		}
	}

	vibesEngine.Logger.Fyi("go-vibes engine started with extra vibes ✨")

	return vibesEngine
}

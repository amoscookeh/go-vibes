package vibes

import (
	"github.com/gin-gonic/gin"
)

func init() {
	SetReleaseMode()
}

// HandlerFunc defines the handler used by vibes middleware as return value
type HandlerFunc func(*Context)

// VibesEngine represents the framework's instance
type VibesEngine struct {
	*gin.Engine
	Logger *VibeLogger
}

// disables Gin's debug logs
func SetReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

// New creates a new VibesEngine instance with minimal middleware
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

// Default creates a new VibesEngine with the default middleware
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

// wrapHandler wraps a Vibes HandlerFunc to a Gin HandlerFunc
func wrapHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(&Context{Context: c})
	}
}

// GET is a shortcut for engine.Handle("GET", path, handle)
func (e *VibesEngine) GET(relativePath string, handler HandlerFunc) {
	e.Engine.GET(relativePath, wrapHandler(handler))
}

// POST is a shortcut for engine.Handle("POST", path, handle)
func (e *VibesEngine) POST(relativePath string, handler HandlerFunc) {
	e.Engine.POST(relativePath, wrapHandler(handler))
}

// PUT is a shortcut for engine.Handle("PUT", path, handle)
func (e *VibesEngine) PUT(relativePath string, handler HandlerFunc) {
	e.Engine.PUT(relativePath, wrapHandler(handler))
}

// DELETE is a shortcut for engine.Handle("DELETE", path, handle)
func (e *VibesEngine) DELETE(relativePath string, handler HandlerFunc) {
	e.Engine.DELETE(relativePath, wrapHandler(handler))
}

// PATCH is a shortcut for engine.Handle("PATCH", path, handle)
func (e *VibesEngine) PATCH(relativePath string, handler HandlerFunc) {
	e.Engine.PATCH(relativePath, wrapHandler(handler))
}

// HEAD is a shortcut for engine.Handle("HEAD", path, handle)
func (e *VibesEngine) HEAD(relativePath string, handler HandlerFunc) {
	e.Engine.HEAD(relativePath, wrapHandler(handler))
}

// OPTIONS is a shortcut for engine.Handle("OPTIONS", path, handle)
func (e *VibesEngine) OPTIONS(relativePath string, handler HandlerFunc) {
	e.Engine.OPTIONS(relativePath, wrapHandler(handler))
}

// Any registers a route that matches all HTTP methods
func (e *VibesEngine) Any(relativePath string, handler HandlerFunc) {
	e.Engine.Any(relativePath, wrapHandler(handler))
}

// Run attaches the router to a http.Server and starts listening
func (e *VibesEngine) Run(addr ...string) error {
	return e.Engine.Run(addr...)
}

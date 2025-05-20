package vibes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Context wraps gin.Context to hide Gin dependency from users
type Context struct {
	*gin.Context
}

// Map is an alias for gin.H to avoid users needing to import gin
type Map map[string]interface{}

// JSON renders the data as JSON with the given status code
func (c *Context) JSON(code int, obj interface{}) {
	c.Context.JSON(code, obj)
}

// String renders the given string with the specified status code
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Context.String(code, format, values...)
}

// HTML renders the string as HTML with the given status code
func (c *Context) HTML(code int, html string) {
	c.Header("Content-Type", "text/html")
	c.String(code, html)
}

// Data renders the raw data with the specified status code
func (c *Context) Data(code int, contentType string, data []byte) {
	c.Context.Data(code, contentType, data)
}

// File serves a file with the given filepath
func (c *Context) File(filepath string) {
	c.Context.File(filepath)
}

// Header sets a response header
func (c *Context) Header(key, value string) {
	c.Context.Header(key, value)
}

// GetHeader returns the value of a request header
func (c *Context) GetHeader(key string) string {
	return c.Context.GetHeader(key)
}

// Param returns the value of a URL parameter
func (c *Context) Param(key string) string {
	return c.Context.Param(key)
}

// Query returns the value of a query parameter
func (c *Context) Query(key string) string {
	return c.Context.Query(key)
}

// DefaultQuery returns the value of a query parameter or the default value
func (c *Context) DefaultQuery(key, defaultValue string) string {
	return c.Context.DefaultQuery(key, defaultValue)
}

// FormValue returns the value of a form field
func (c *Context) FormValue(key string) string {
	return c.Context.PostForm(key)
}

// Bind binds the request body to the provided struct
func (c *Context) Bind(obj interface{}) error {
	return c.Context.ShouldBind(obj)
}

// BindJSON binds the JSON request body to the provided struct
func (c *Context) BindJSON(obj interface{}) error {
	return c.Context.ShouldBindJSON(obj)
}

// StatusCodes provides constants to avoid importing http package
var StatusCodes = struct {
	OK                  int
	Created             int
	NoContent           int
	BadRequest          int
	Unauthorized        int
	Forbidden           int
	NotFound            int
	MethodNotAllowed    int
	Conflict            int
	InternalServerError int
}{
	OK:                  http.StatusOK,
	Created:             http.StatusCreated,
	NoContent:           http.StatusNoContent,
	BadRequest:          http.StatusBadRequest,
	Unauthorized:        http.StatusUnauthorized,
	Forbidden:           http.StatusForbidden,
	NotFound:            http.StatusNotFound,
	MethodNotAllowed:    http.StatusMethodNotAllowed,
	Conflict:            http.StatusConflict,
	InternalServerError: http.StatusInternalServerError,
}

package main

import (
	"net/http"

	"github.com/amoscookeh/go-vibes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new vibes engine
	r := vibes.Default()

	// Define routes
	r.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All good! Check the status emoji!",
		})
	})

	r.GET("/created", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Resource created! Check the status emoji!",
		})
	})

	r.GET("/bad-request", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request! Check the status emoji!",
		})
	})

	r.GET("/not-found", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found! Check the status emoji!",
		})
	})

	r.GET("/server-error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server error! Check the status emoji!",
		})
	})

	// HTML response example
	r.GET("/html-ok", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, `
			<html>
				<head>
					<title>Vibes Framework Demo</title>
					<style>
						body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
						h1 { color: #333; }
						.container { background: #f9f9f9; border-radius: 10px; padding: 20px; }
					</style>
				</head>
				<body>
					<div class="container">
						<h1>Vibes Framework HTML Demo</h1>
						<p>This is a successful response! Look for the emoji status badge in the corner.</p>
						<p>The actual HTTP status code is always 200, but the true status is communicated via emojis.</p>
						<p><a href="/html-error">See an error page</a></p>
					</div>
				</body>
			</html>
		`)
	})

	// HTML error example
	r.GET("/html-error", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusNotFound, `
			<html>
				<head>
					<title>Vibes Framework Demo - Error</title>
					<style>
						body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
						h1 { color: #d9534f; }
						.container { background: #f9f9f9; border-radius: 10px; padding: 20px; }
					</style>
				</head>
				<body>
					<div class="container">
						<h1>Not Found Error!</h1>
						<p>This page demonstrates a 404 Not Found error response.</p>
						<p>Look for the emoji status badge in the corner!</p>
						<p><a href="/html-ok">Back to success page</a></p>
					</div>
				</body>
			</html>
		`)
	})

	// Plain text response example
	r.GET("/text", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		c.String(http.StatusOK, "This is a plain text response. The emoji status should be prepended.")
	})

	// Run the server
	r.Run(":8080")
}

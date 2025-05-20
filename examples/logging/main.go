package main

import (
	"net/http"
	"os"

	"github.com/amoscookeh/go-vibes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := vibes.Default()

	logger := r.Logger

	logger.Fyi("Server is starting up! Just FYI!")
	logger.Uhoh("Memory usage seems a bit high, but we'll manage!")

	// successful route (will use FYI log level)
	r.GET("/success", func(c *gin.Context) {
		logger.Fyi("Processing success request")
		c.JSON(http.StatusOK, gin.H{
			"message": "This is a success!",
		})
	})

	// not found route (will use UHOH log level)
	r.GET("/not-found", func(c *gin.Context) {
		logger.Uhoh("Someone tried to access something that doesn't exist!")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can't find what you're looking for!",
		})
	})

	// error route (will use CRAP log level)
	r.GET("/error", func(c *gin.Context) {
		logger.Crap("Something bad happened in the server!")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server is having issues!",
		})
	})

	// fatal route (will use COOKED log level and exit)
	r.GET("/fatal", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "This would normally crash the server!",
		})

		go func() {
			logger.Uhoh("About to demonstrate a COOKED log (fatal)...")
			logger.Cooked("The server is completely COOKED! Shutting down!")
		}()
	})

	logger.Fyi("Server is ready on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Crap("Failed to start server: %v", err)
		os.Exit(1)
	}
}

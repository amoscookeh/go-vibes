package main

import (
	"os"

	"github.com/amoscookeh/go-vibes"
)

func main() {
	r := vibes.Default()

	logger := r.Logger

	logger.Fyi("Server is starting up! Just FYI!")
	logger.Uhoh("Memory usage seems a bit high, but we'll manage!")

	// successful route (will use FYI log level)
	r.VIBE("/success", func(c *vibes.Context) {
		logger.Fyi("Processing success request")
		c.JSON(vibes.StatusCodes.OK, vibes.Map{
			"message": "This is a success!",
		})
	})

	// not found route (will use UHOH log level)
	r.VIBE("/not-found", func(c *vibes.Context) {
		logger.Uhoh("Someone tried to access something that doesn't exist!")
		c.JSON(vibes.StatusCodes.NotFound, vibes.Map{
			"message": "Can't find what you're looking for!",
		})
	})

	// error route (will use CRAP log level)
	r.VIBE("/error", func(c *vibes.Context) {
		logger.Crap("Something bad happened in the server!")
		c.JSON(vibes.StatusCodes.InternalServerError, vibes.Map{
			"message": "Server is having issues!",
		})
	})

	// fatal route (will use COOKED log level and exit)
	r.VIBE("/fatal", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.InternalServerError, vibes.Map{
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

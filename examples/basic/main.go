package main

import (
	"github.com/amoscookeh/go-vibes"
)

func main() {
	// Create a new vibes engine
	r := vibes.Default()

	// Define routes
	r.VIBE("/ok", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.OK, vibes.Map{
			"message": "All good! Check the status emoji!",
		})
	})

	r.VIBE("/created", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.Created, vibes.Map{
			"message": "Resource created! Check the status emoji!",
		})
	})

	r.VIBE("/bad-request", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.BadRequest, vibes.Map{
			"message": "Bad request! Check the status emoji!",
		})
	})

	r.VIBE("/not-found", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.NotFound, vibes.Map{
			"message": "Not found! Check the status emoji!",
		})
	})

	r.VIBE("/server-error", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.InternalServerError, vibes.Map{
			"message": "Server error! Check the status emoji!",
		})
	})

	// HTML response example
	r.VIBE("/html-ok", func(c *vibes.Context) {
		c.HTML(vibes.StatusCodes.OK, `
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
	r.VIBE("/html-error", func(c *vibes.Context) {
		c.HTML(vibes.StatusCodes.NotFound, `
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

	r.VIBE("/text", func(c *vibes.Context) {
		c.String(vibes.StatusCodes.OK, "This is a plain text response. The emoji status should be prepended.")
	})

	r.Run(":8080")
}

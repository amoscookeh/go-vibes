# 🚀 Quick Start

_Give your API endpoints feelings in 5 minutes or less_

## Your First Vibey Application

If you haven't installed Go-Vibes yet, check out the [Installation Guide](installation.md). Already
vibey? Let's proceed!

### Step 1: Create a new project

Create a directory for your emotionally intelligent application:

```bash
mkdir my-vibey-api
cd my-vibey-api
```

Initialize your Go module:

```bash
go mod init github.com/yourusername/my-vibey-api
```

### Step 2: Install Go-Vibes

```bash
go get github.com/amoscookeh/go-vibes
```

### Step 3: Create your first vibey API

Create a file named `main.go` with this emotionally expressive code:

```go
package main

import (
	"github.com/amoscookeh/go-vibes"
)

func main() {
	// Create a new vibey engine
	r := vibes.Default()

	// Get direct access to the emotional logger
	logger := r.Logger
	logger.Fyi("Starting my first vibey application! SO EXCITED!")

	// Define some routes with emotional responses
	r.VIBE("/happy", func(c *vibes.Context) {
		logger.Fyi("Someone is looking for happiness!")
		c.JSON(vibes.StatusCodes.OK, vibes.Map{
			"mood": "ecstatic",
			"message": "Everything is awesome!",
		})
		// Automatically adds ✅👌🆗 to your response!
	})

	r.VIBE("/sad", func(c *vibes.Context) {
		logger.Uhoh("Oh no, someone is feeling blue...")
		c.JSON(vibes.StatusCodes.BadRequest, vibes.Map{
			"mood": "melancholy",
			"message": "Sorry you're feeling down",
		})
		// Automatically adds 💔👿😭 to your response!
	})

	r.VIBE("/lost", func(c *vibes.Context) {
		logger.Crap("Someone got lost in our API!")
		c.JSON(vibes.StatusCodes.NotFound, vibes.Map{
			"mood": "confused",
			"message": "This pathway leads nowhere",
		})
		// Automatically adds 🕵️🔍❓ to your response!
	})

	// Run the server with good vibes only
	logger.Fyi("Server is vibing on http://localhost:8080")
	r.Run(":8080")
}
```

### Step 4: Run your application

```bash
go run main.go
```

You should see your emotional logger spring to life:

```
[15:04:05] FYI 💁: Starting my first vibey application! SO EXCITED!
[15:04:05] FYI 💁: Server is vibing on http://localhost:8080
```

### Step 5: Test your vibey API

Use curl or your browser to experience the emotional responses:

```bash
# Happy route - using vibey HTTP method
curl -X VIBE http://localhost:8080/happy

# Standard HTTP method also works through automatic conversion
curl http://localhost:8080/happy

# Response:
# {"mood":"ecstatic","message":"Everything is awesome!","status_emoji":"✅👌🆗"}
```

```bash
# Sad route - using vibey HTTP method
curl -X VIBE http://localhost:8080/sad

# Response:
# {"mood":"melancholy","message":"Sorry you're feeling down","status_emoji":"💔👿😭"}
```

```bash
# Lost route - using vibey HTTP method
curl -X VIBE http://localhost:8080/lost

# Response:
# {"mood":"confused","message":"This pathway leads nowhere","status_emoji":"🕵️🔍❓"}
```

## What Just Happened?

Congratulations! You've just created an API that:

1. Uses vibey HTTP methods (VIBE, MANIFEST, ALIGN, RELEASE) instead of boring standard methods
2. Communicates using emoji status codes instead of boring numbers
3. Logs with emotional expression instead of severity levels
4. Has more personality than 99% of APIs in production today

## Next Steps

Now that you've created your first vibey application, check out:

- [Configuration Guide](configuration.md) to customize your vibes
- [Emoji Status Responses](../guides/emoji-status.md) for detailed information on status emojis
- [Emotional Logging](../guides/emotional-logging.md) to learn the full range of emotional
  expressions

## Troubleshooting

**Problem**: My application shows no emotions.  
**Solution**: Check that you're using `vibes.Default()` and not plain `gin.Default()`. Don't
suppress your code's feelings!

**Problem**: My team doesn't appreciate my vibey code.  
**Solution**: Find a new team. Life's too short to work with people who don't appreciate good vibes.

**Problem**: I can't stop adding emojis to everything now.  
**Solution**: This is not a problem. This is enlightenment. 🧘‍♂️

**Problem**: My application panics when I use GET, POST, PUT, or DELETE.  
**Solution**: Use the vibey alternatives: VIBE, MANIFEST, ALIGN, and RELEASE. Your code will thank
you for the positive energy flow!

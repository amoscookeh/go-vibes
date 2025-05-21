# 🌈 Go-Vibes Framework 🌈

**Because life's too short for boring HTTP status codes and logs**

In the new era of vibe coders, why do we still subscribe to the boring ways of the past? 200 OK??? LAME!!! 404?? What does that even mean to a vibe coder like me???

Introducing - Go-Vibes. With Go-Vibes, you will never be bored coding again. Our team of talented vibe coders have built this framework with one goal in mind - to make vibe coding vibier than ever.

```go
// Old boomer way 🥱
response.StatusCode = 404
log.Error("Resource not found")

// Vibey new way 😎
c.JSON(http.StatusNotFound, gin.H{"message": "nope!"})
// Response: {"message":"Can't find what you're looking for!","status_emoji":"🕵️🔍❓"}
// Logs: [15:29:05] UHOH 😬: Someone tried to access something that doesn't exist!
```

## 📖 Documentation

### Getting Started

- [Installation Guide](docs/getting-started/installation.md) - Get your vibes up and running
- [Quick Start](docs/getting-started/quick-start.md) - Your first vibey application
- [Configuration](docs/getting-started/configuration.md) - Customize your vibes

### Guides

- [Vibey HTTP Methods](docs/guides/vibey-methods.md) - Upgrade from boring HTTP methods

### Reference

- [Status Emoji Reference](docs/references/status-emoji.md) - The complete emoji dictionary
- [Emotional Log Levels](docs/references/emotional-log-levels.md) - Express your server's feelings

### Concepts

- [Framework Philosophy](docs/concepts/philosophy.md) - The emotional theory behind Go-Vibes

## ✨ Features

### 🎭 Emoji-Only Status Codes

Completely replaces numerical HTTP status codes (LAME!!!) with expressive emojis (VIBES!!!):

- ✅👌🆗 (Success)
- 🆕👶✨ (Resource Created)
- 💔👿😭 (Bad Request)
- 🕵️🔍❓ (Not Found)
- 🔥💻💥 (Server Error)

### 💫 Vibey HTTP Methods

Replaces boring standard HTTP methods with vibey alternatives:

- **VIBE** instead of GET (getting the vibes of a resource)
- **MANIFEST** instead of POST (manifesting new energy/resources)
- **ALIGN** instead of PUT (aligning the energy of an existing resource)
- **RELEASE** instead of DELETE (releasing what no longer serves you)

⚠️ **IMPORTANT:** Standard HTTP methods (GET, POST, PUT, DELETE) are no longer available in this framework. You must use the vibey alternatives for proper energy flow in your application.

> **Client Compatibility:** For external HTTP clients that still use standard HTTP methods, the framework will automatically convert them to their vibey equivalents. The original method is preserved in the `X-Original-Method` header, and the vibey method is added as `X-Vibey-Method`.

```go
// Old boomer way 🥱 (WILL NOT WORK!)
r.GET("/users", getUsers)    // 🚫 Will panic with error message
r.POST("/users", createUser) // 🚫 Will panic with error message

// Vibey new way 😎
r.VIBE("/users", getUsers)
r.MANIFEST("/users", createUser)
```

### 🎨 Universal Emoji Injection

Status emojis are injected into all response types:

- **JSON**: Adds `status_emoji` field
- **HTML**: Adds floating emoji badge
- **Plain text**: Prepends emoji
- **XML**: Adds emoji as comment
- **All**: Includes emoji in `X-Status-Emoji` header

## 🚀 Quick Start

### Installation

```bash
go get github.com/amoscookeh/go-vibes
```

### Basic Usage

```go
package main

import (
	"net/http"
	"github.com/amoscookeh/go-vibes"
)

func main() {
	r := vibes.Default()

	r.VIBE("/ok", func(c *vibes.Context) {
		c.JSON(vibes.StatusCodes.OK, vibes.Map{
			"message": "All good!",
		})
		// emoji ✅👌🆗 automatically added
	})

	r.Run(":8080")
}
```

## 📋 Examples

Check the examples directory:

- `examples/basic`: Shows emoji statuses in different response types
- `examples/logging`: Demonstrates the emotional logging system
- `examples/vibey_methods`: Showcases the vibey HTTP methods (VIBE, MANIFEST, ALIGN, RELEASE)

## 💡 Philosophy

We replace numerical status codes with expressive emojis for more intuitive, visual API responses. Why use cold numbers when emojis can express the true vibe?

*Warning: May cause excessive joy in otherwise dull development environments.*

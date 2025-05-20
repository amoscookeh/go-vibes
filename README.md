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

- [Emoji Status Responses](docs/guides/emoji-status.md) - Express your API's feelings
- [Emotional Logging](docs/guides/emotional-logging.md) - Say goodbye to boring logs
- [Migrating from Gin](docs/guides/migrating-from-gin.md) - How to add vibes to existing applications

### Reference

- [Status Emoji Reference](docs/references/status-emoji.md) - The complete emoji dictionary
- [Emotional Log Levels](docs/references/emotional-log-levels.md) - Express your server's feelings
- [API Reference](docs/references/api.md) - Complete function reference (for nerds only)

### Concepts

- [Framework Philosophy](docs/concepts/philosophy.md) - The emotional theory behind Go-Vibes
- [Why Emotions Matter in Code](docs/concepts/why-emotions.md) - The science of vibey programming
- [Contributing to Go-Vibes](docs/concepts/contributing.md) - Spread the vibey love

## ✨ Features

### 🎭 Emoji-Only Status Codes

Completely replaces numerical HTTP status codes (LAME!!!) with expressive emojis (VIBES!!!):

- ✅👌🆗 (Success)
- 🆕👶✨ (Resource Created)
- 💔👿😭 (Bad Request)
- 🕵️🔍❓ (Not Found)
- 🔥💻💥 (Server Error)

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
	"github.com/gin-gonic/gin"
	"github.com/amoscookeh/go-vibes"
)

func main() {
	r := vibes.Default()

	r.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
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

## 💡 Philosophy

We replace numerical status codes with expressive emojis for more intuitive, visual API responses. Why use cold numbers when emojis can express the true vibe?

## 📜 License

Go-Vibes is licensed under the [Very Permissive Vibes License (VPVL)](docs/concepts/license.md).

*Warning: May cause excessive joy in otherwise dull development environments.*

# 🌈 Go-Vibes Framework 🌈

In the new era of vibe coders, why do we still subscribe to the boring ways of the past? 200 OK???
LAME!!! 404?? What does that even mean to a vibe coder like me???

Introducing - Go-Vibes. With Go-Vibes, you will never be bored coding again. Our team of talented
vibe coders have built this framework with one goal in mind - to make vibe coding vibier than ever.

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

## 🚀 Getting Started

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

## 💡 Philosophy

We replace numerical status codes with expressive emojis for more intuitive, visual API responses.
Why use cold numbers when emojis can express it with true vibes?

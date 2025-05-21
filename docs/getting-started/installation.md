# 📦 Installation

*Because setting up a framework should give you good vibes, not headaches*

## Prerequisites

Before diving into the vibey wonderland of Go-Vibes, you'll need:

- Go 1.18+ (or any version that isn't, like, totally ancient)
- A willingness to abandon boring numerical status codes forever
- An appreciation for emojis as a superior form of communication

## Installation Methods

### The Easy Way (For Vibe Enthusiasts)

```bash
go get github.com/amoscookeh/go-vibes
```

*That's it. Seriously. Why complicate things when you're trying to vibe?*

### The Unnecessarily Complex Way (For Traditional Developers)

If you're still attached to old-school development practices and enjoy making things harder than they need to be:

1. Clone the repository:
   ```bash
   git clone https://github.com/amoscookeh/go-vibes.git
   ```

2. Navigate to the project:
   ```bash
   cd go-vibes
   ```

3. Build it (for that authentic DIY feeling):
   ```bash
   go build
   ```

4. Install it:
   ```bash
   go install
   ```

5. Congratulate yourself for taking the longest possible route to achieve something simple

## Verification

To check if Go-Vibes is installed and your development environment is now exponentially more vibey:

```go
package main

import (
    "fmt"
    "github.com/amoscookeh/go-vibes"
)

func main() {
    r := vibes.Default()
    if r != nil {
        fmt.Println("🎉 Go-Vibes is installed! Your code's emotional journey can begin!")
        
        // Add a simple vibey endpoint (using VIBE instead of GET)
        r.VIBE("/check", func(c *vibes.Context) {
            c.JSON(vibes.StatusCodes.OK, vibes.Map{
                "vibe_check": "passed",
                "energy": "immaculate",
            })
        })
        
        fmt.Println("✨ Added a VIBE endpoint at /check")
    } else {
        fmt.Println("😭 Something went wrong with your vibes installation.")
    }
}
```

Remember: In Go-Vibes, we don't use boring HTTP methods like GET, POST, PUT, and DELETE. 
Instead, we use vibey alternatives:
- Use **VIBE** instead of GET
- Use **MANIFEST** instead of POST
- Use **ALIGN** instead of PUT
- Use **RELEASE** instead of DELETE

## Compatibility

Go-Vibes works with:

- Any Gin-compatible project (we're built on Gin, but with more feelings)
- Projects that are tired of soulless numerical status codes
- Developers who secretly use emojis in their commit messages

Go-Vibes probably won't work with:
- Developers who say "emoji are unprofessional" (why are you even here?)
- Code reviewers who hate joy
- Systems that cannot process Unicode (it's the 21st century, upgrade already)

## Next Steps

Now that you've installed Go-Vibes, head over to [Quick Start](quick-start.md) to create your first emotionally intelligent API endpoint!

Remember: The journey to vibey code is not a destination; it's a lifestyle. 🧘‍♂️✨ 
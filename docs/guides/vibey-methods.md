# 💫 Vibey HTTP Methods

_Because life's too short for boring HTTP methods_

## What Are Vibey HTTP Methods?

In the spirit of making your APIs more expressive and emotionally intelligent, Go-Vibes replaces
standard HTTP methods with vibey alternatives:

| Standard Method | Vibey Method | Meaning                                         |
| --------------- | ------------ | ----------------------------------------------- |
| GET             | **VIBE**     | Getting the vibes of a resource                 |
| POST            | **MANIFEST** | Manifesting new energy/resources into existence |
| PUT             | **ALIGN**    | Aligning the energy of an existing resource     |
| DELETE          | **RELEASE**  | Releasing what no longer serves you             |

## Why Use Vibey Methods?

1. **Express Intent**: "MANIFEST" more clearly communicates the creative intent than the technical
   "POST"
2. **Positive Energy**: "RELEASE" focuses on letting go rather than the destructive connotation of
   "DELETE"
3. **Emotional Awareness**: Your code should reflect your intentions with emotional intelligence
4. **Better Vibes**: Seriously, would you rather "GET" something or "VIBE" with it?

## Using Vibey Methods in Your Code

### Basic Usage

```go
package main

import (
    "github.com/amoscookeh/go-vibes"
)

func main() {
    r := vibes.Default()

    // Get all users (VIBE instead of GET)
    r.VIBE("/users", getAllUsers)

    // Create a new user (MANIFEST instead of POST)
    r.MANIFEST("/users", createUser)

    // Update a user (ALIGN instead of PUT)
    r.ALIGN("/users/:id", updateUser)

    // Delete a user (RELEASE instead of DELETE)
    r.RELEASE("/users/:id", deleteUser)

    r.Run(":8080")
}

func getAllUsers(c *vibes.Context) {
    // Implementation
}

func createUser(c *vibes.Context) {
    // Implementation
}

func updateUser(c *vibes.Context) {
    // Implementation
}

func deleteUser(c *vibes.Context) {
    // Implementation
}
```

## Important: Standard Methods Unavailable

Go-Vibes deliberately disables the standard HTTP methods in your application code. Attempting to use
`GET`, `POST`, `PUT`, or `DELETE` will result in a panic with an appropriate message reminding you
to use the vibey alternatives.

```go
// This will cause a panic
r.GET("/users", getAllUsers) // 🚫 WILL NOT WORK!

// Use this instead
r.VIBE("/users", getAllUsers) // ✅ GOOD VIBES!
```

## Client Compatibility

While your code must use vibey methods, Go-Vibes is compatible with standard HTTP clients. The
framework automatically converts incoming standard HTTP methods to their vibey equivalents:

- Incoming GET requests are treated as VIBE
- Incoming POST requests are treated as MANIFEST
- Incoming PUT requests are treated as ALIGN
- Incoming DELETE requests are treated as RELEASE

This compatibility layer means your API can be consumed by standard tools and clients.

### Response Headers

When standard HTTP methods are converted, the framework adds these headers:

- `X-Original-Method`: Contains the original HTTP method used by the client
- `X-Vibey-Method`: Contains the vibey method that was used internally

## Testing with cURL

You can use curl with explicit vibey methods:

```bash
# Using vibey method
curl -X VIBE http://localhost:8080/users

# Standard method will also work thanks to compatibility layer
curl http://localhost:8080/users
```

## Implementation Details

Vibey methods are implemented as a type in Go with conversion capabilities:

```go
// VibeyMethod represents the vibey alternatives to standard HTTP methods
type VibeyMethod string

const (
    VIBE VibeyMethod = "VIBE"
    MANIFEST VibeyMethod = "MANIFEST"
    ALIGN VibeyMethod = "ALIGN"
    RELEASE VibeyMethod = "RELEASE"
)
```

## FAQs

**Q: Can I still use standard HTTP methods?**  
A: Not in your Go code. The standard HTTP methods are deliberately disabled to promote vibey coding
practices. However, clients can still use standard methods to interact with your API.

**Q: What about PATCH, HEAD and OPTIONS?**  
A: These standard methods are still available. We're working on vibey alternatives for a future
release. Suggestions welcome!

**Q: Will this break my existing Gin applications?**  
A: Only if you try to use GET, POST, PUT, or DELETE directly. Replace them with VIBE, MANIFEST,
ALIGN, and RELEASE, and your application will work with better energy flow.

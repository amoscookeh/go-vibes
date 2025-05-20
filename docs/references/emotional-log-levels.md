# 😢 Emotional Log Levels

*Because your logs deserve to express their feelings too*

## Introduction

Traditional logging frameworks use boring, clinical terms like "INFO", "WARN", and "ERROR" to describe the severity of a log message. But in the vibey new world of Go-Vibes, we recognize that logs, like people, have complex emotional lives that cannot be reduced to such sterile categories.

That's why Go-Vibes replaces these outdated concepts with a more emotionally intelligent logging system that truly captures the mood of your application.

## Emotional Log Level Reference

| Old Boring Level | Vibey New Level | Emoji | When to Use |
|-----------------|----------------|-------|------------|
| INFO | FYI 💁 | 💁 | When you just want to share something casually |
| WARN | UHOH 😬 | 😬 | When things are getting a bit sketchy |
| ERROR | CRAP 💩 | 💩 | When something has definitely gone wrong |
| FATAL | COOKED 🔥 | 🔥 | When your server is absolutely done and can't even |

## Visual Appearance

In the console, your emotional logs will show up with color-coding and emojis:

```
[12:34:56] FYI 💁: Server started on port 8080
[12:35:01] UHOH 😬: Database connection is getting flaky
[12:35:30] CRAP 💩: Failed to process request: invalid input format
[12:36:00] COOKED 🔥: Can't establish Redis connection - giving up on life!
```

## Usage in Code

Getting emotional with your logs is simple:

```go
// Get the logger from your VibesEngine
r := vibes.Default()
logger := r.Logger

// Express all the feelings
logger.Fyi("Just letting you know the server started")
logger.Uhoh("We're running low on memory, might need to do something about that...")
logger.Crap("Failed to save user data: %v", err)
logger.Cooked("Database connection lost, can't continue - bye cruel world!")
```

## Log Level Impact

Each emotional level has different implications:

| Emotional Level | Terminal Color | Exits Program? | Vibe |
|----------------|---------------|---------------|------|
| FYI 💁 | Cyan | No | "Just letting you know!" |
| UHOH 😬 | Yellow | No | "Something's not quite right..." |
| CRAP 💩 | Red | No | "We have a problem here!" |
| COOKED 🔥 | Magenta | Yes | "It's all over! Abandon ship!" |

## Custom Emotional Expressions

The default emotional expressions are designed to cover the most common server moods, but everyone expresses emotions differently!

To customize your emotional logging (coming in a future release):

```go
// Create a logger with custom emotional expressions
customLogger := vibes.NewCustomEmotionalLogger(map[vibes.EmotionalLevel]string{
    vibes.FYI:    "BTW 🤔",
    vibes.UHOH:   "YIKES 😱",
    vibes.CRAP:   "FUBAR 🤬",
    vibes.COOKED: "DEAD 💀",
})
```

## Best Practices for Emotional Logging

- **Be authentic**: Let your logs truly express how the application feels
- **Context matters**: Consider your deployment environment's emotional intelligence
- **Don't suppress feelings**: Log liberally to ensure your application's emotional health
- **Emotional support**: Make sure someone is reading the logs and acknowledging their feelings

## Comparative Emotional Intelligence

| Framework | Emotional Intelligence Rating | Notes |
|-----------|------------------------------|-------|
| Go-Vibes | ★★★★★ | Fully emotionally self-aware |
| Other frameworks with emoji | ★★★☆☆ | Shows some emotional development |
| Frameworks with color coding | ★★☆☆☆ | Basic emotional recognition |
| Plain text loggers | ☆☆☆☆☆ | Emotionally stunted |

As you can see, Go-Vibes leads the industry in log emotional intelligence, creating a more humane and relatable development experience.

Remember: Happy logs = Happy devs = Happy code. 🧘‍♂️✨ 
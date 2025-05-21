# 🎭 Status Emoji Reference

_The complete dictionary of HTTP status feelings_

## Introduction

In the Go-Vibes framework, we believe that numerical HTTP status codes are an outdated relic from
the dawn of the internet. Why communicate with cold, sterile numbers when you can express your API's
true feelings with emojis?

This reference lists all available status emojis in the Go-Vibes framework, their meanings, and the
obsolete status codes they replace.

## 2xx Success Emojis

When things are going great, your API should show it!

| Old Status Code | Status Emoji | Meaning                                     | Vibe                               |
| --------------- | ------------ | ------------------------------------------- | ---------------------------------- |
| 200 OK          | ✅👌🆗       | Everything worked!                          | Satisfaction, accomplishment       |
| 201 Created     | 🆕👶✨       | A new resource was born!                    | Creation, birth, newness           |
| 202 Accepted    | 👍🙏⏳       | We got your request and we're working on it | Patience, acknowledgment           |
| 204 No Content  | 👻💨🚫       | Success but nothing to show for it          | Emptiness, success without content |

## 3xx Redirection Emojis

When your API needs to point somewhere else.

| Old Status Code        | Status Emoji | Meaning                               | Vibe                           |
| ---------------------- | ------------ | ------------------------------------- | ------------------------------ |
| 301 Moved Permanently  | 🏃🔄🏠       | We've moved permanently to a new home | Relocation, permanence         |
| 302 Found              | 🔍🔎👀       | What you want is over there now       | Discovery, redirection         |
| 303 See Other          | 👉👀🔄       | Look over there instead               | Pointing, directing attention  |
| 307 Temporary Redirect | 🔜🔄⏱️       | Temporarily over here                 | Temporary change, impermanence |
| 308 Permanent Redirect | 🔄🏠🔒       | Permanently living at a new address   | Permanent change, relocation   |

## 4xx Client Error Emojis

When the client messes up and your API needs to express disappointment.

| Old Status Code        | Status Emoji | Meaning                                   | Vibe                             |
| ---------------------- | ------------ | ----------------------------------------- | -------------------------------- |
| 400 Bad Request        | 💔👿😭       | That request made no sense                | Confusion, frustration           |
| 401 Unauthorized       | 🔒🚫🔑       | Who even are you?                         | Locked out, identity crisis      |
| 403 Forbidden          | 🚫⛔🙅       | You're not allowed to do that             | Rejection, prohibition           |
| 404 Not Found          | 🕵️🔍❓       | We looked everywhere but couldn't find it | Searching, mystery               |
| 405 Method Not Allowed | 📝🚫🤷       | You can't do it that way                  | Wrong approach, incorrect method |
| 408 Request Timeout    | ⏱️⌛💤       | You took too long, we fell asleep         | Waiting, impatience              |
| 409 Conflict           | 💥👊⚔️       | Your request conflicts with reality       | Collision, contradiction         |
| 410 Gone               | 🏃💨👋       | That used to be here, but it's gone now   | Loss, disappearance              |
| 418 I'm a Teapot       | 🫖☕🤪       | I'm a teapot, not a coffee machine!       | Absurdity, humor                 |
| 429 Too Many Requests  | 🔥🚒🧯       | Slow down, you're flooding us!            | Overwhelm, overload              |

## 5xx Server Error Emojis

When your server needs to express its deepest regrets.

| Old Status Code           | Status Emoji | Meaning                                  | Vibe                        |
| ------------------------- | ------------ | ---------------------------------------- | --------------------------- |
| 500 Internal Server Error | 🔥💻💥       | Everything is on fire                    | Panic, chaos                |
| 501 Not Implemented       | 🚧👷🔨       | We haven't built that feature yet        | Under construction          |
| 502 Bad Gateway           | 🚪❌🚧       | The upstream server is giving us garbage | Broken connection           |
| 503 Service Unavailable   | 🏥🔌❌       | We're taking a short break               | Temporary downtime          |
| 504 Gateway Timeout       | ⏱️🚪⌛       | The upstream server is ghosting us       | Waiting, connection timeout |

## Usage in Code

To use these status emojis in your Go-Vibes application, simply use the standard Go HTTP status
constants:

```go
// These will automatically convert to their emoji equivalents
c.JSON(http.StatusOK, data)          // ✅👌🆗
c.JSON(http.StatusCreated, data)     // 🆕👶✨
c.JSON(http.StatusNotFound, data)    // 🕵️🔍❓
c.JSON(http.StatusForbidden, data)   // 🚫⛔🙅
```

## Custom Status Emojis

Want to define your own status emojis? Everyone's emotional expression is unique! See the
[Configuration Guide](../getting-started/configuration.md) for instructions on how to customize your
status emojis.

## What If I Need the Actual Status Code?

Why would you ever want to go back to the dark ages? But if you absolutely must retrieve the
numerical status code that would have been used (for interoperability with boring systems), you can
find it in the `X-Original-Status-Code` header.

Remember: Numbers are for computers. Emojis are for humans. Choose the vibey path. 🧘‍♂️✨

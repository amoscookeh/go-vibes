package vibes

import (
	"net/http"
)

var StatusEmojiMap = map[int]string{
	http.StatusOK:                  "✅👌🆗",  // 200
	http.StatusCreated:             "🆕👶✨",  // 201
	http.StatusAccepted:            "👍🙏⏳",  // 202
	http.StatusNoContent:           "👻💨🚫",  // 204
	http.StatusMovedPermanently:    "🏃🔄🏠",  // 301
	http.StatusFound:               "🔍🔎👀",  // 302
	http.StatusSeeOther:            "👉👀🔄",  // 303
	http.StatusBadRequest:          "💔👿😭",  // 400
	http.StatusUnauthorized:        "🔒🚫🔑",  // 401
	http.StatusForbidden:           "🚫⛔🙅",  // 403
	http.StatusNotFound:            "🕵️🔍❓", // 404
	http.StatusMethodNotAllowed:    "📝🚫🤷",  // 405
	http.StatusRequestTimeout:      "⏱️⌛💤", // 408
	http.StatusConflict:            "💥👊⚔️", // 409
	http.StatusGone:                "🏃💨👋",  // 410
	http.StatusInternalServerError: "🔥💻💥",  // 500
	http.StatusNotImplemented:      "🚧👷🔨",  // 501
	http.StatusBadGateway:          "🚪❌🚧",  // 502
	http.StatusServiceUnavailable:  "🏥🔌❌",  // 503
	http.StatusGatewayTimeout:      "⏱️🚪⌛", // 504
}

func GetStatusEmoji(statusCode int) string {
	if emoji, exists := StatusEmojiMap[statusCode]; exists {
		return emoji
	}
	return "❓❓❓"
}

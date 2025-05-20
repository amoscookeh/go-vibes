package vibes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// VibesResponseWriter wraps the gin.ResponseWriter to inject emoji status
type VibesResponseWriter struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

// NewVibesResponseWriter creates a new VibesResponseWriter
func NewVibesResponseWriter(c *gin.Context) *VibesResponseWriter {
	return &VibesResponseWriter{
		ResponseWriter: c.Writer,
		body:           &bytes.Buffer{},
		statusCode:     http.StatusOK,
	}
}

// WriteHeader overrides the original WriteHeader to capture the status code
func (w *VibesResponseWriter) WriteHeader(code int) {
	w.statusCode = code

	w.Header().Set(XStatusEmoji, GetStatusEmoji(code))
	w.Header().Set("X-Original-Status-Code", "hidden")

	// write 200 OK to hide status_code from clients
	w.ResponseWriter.WriteHeader(http.StatusOK)
}

// Write overrides the original Write to capture the response body
func (w *VibesResponseWriter) Write(data []byte) (int, error) {
	contentType := w.Header().Get("Content-Type")
	emojiStatus := GetStatusEmoji(w.statusCode)

	if strings.Contains(contentType, "application/json") {
		var jsonData map[string]interface{}
		if err := json.Unmarshal(data, &jsonData); err == nil {
			jsonData["status_emoji"] = emojiStatus

			delete(jsonData, "status")
			delete(jsonData, "code")
			delete(jsonData, "statusCode")

			if newData, err := json.Marshal(jsonData); err == nil {
				return w.ResponseWriter.Write(newData)
			}
		}

		return w.ResponseWriter.Write([]byte(fmt.Sprintf(`{"data":%s,"status_emoji":"%s"}`, data, emojiStatus)))
	} else if strings.Contains(contentType, "text/html") {
		badge := fmt.Sprintf(`<div style="position:fixed;bottom:10px;right:10px;background:#f0f0f0;padding:5px;border-radius:5px;font-size:24px;z-index:9999">%s</div>`, emojiStatus)
		htmlWithEmoji := fmt.Sprintf(`<!-- Status Emoji: %s -->%s%s`, emojiStatus, string(data), badge)
		return w.ResponseWriter.Write([]byte(htmlWithEmoji))
	} else if strings.Contains(contentType, "text/plain") {
		textWithEmoji := fmt.Sprintf("Status: %s\n%s", emojiStatus, string(data))
		return w.ResponseWriter.Write([]byte(textWithEmoji))
	} else if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		xmlWithEmoji := fmt.Sprintf(`<!-- Status Emoji: %s -->%s`, emojiStatus, string(data))
		return w.ResponseWriter.Write([]byte(xmlWithEmoji))
	}

	return w.ResponseWriter.Write(data)
}

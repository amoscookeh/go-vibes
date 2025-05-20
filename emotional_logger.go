package vibes

import (
	"fmt"
	"io"
	"os"
	"time"
)

// EmotionalLevel represents logging levels but with vibes
type EmotionalLevel int

const (
	// FYI => INFO
	FYI EmotionalLevel = iota
	// UHOH => WARN
	UHOH
	// CRAP => ERROR
	CRAP
	// COOKED => FATAL
	COOKED
)

var emotionalLevelText = map[EmotionalLevel]string{
	FYI:    "FYI 💁",
	UHOH:   "UHOH 😬",
	CRAP:   "CRAP 💩",
	COOKED: "COOKED 🔥",
}

var emotionalColor = map[EmotionalLevel]string{
	FYI:    "\033[36m", // cyan
	UHOH:   "\033[33m", // yellow
	CRAP:   "\033[31m", // red
	COOKED: "\033[35m", // magenta
}

const resetColor = "\033[0m"

type VibeLogger struct {
	out      io.Writer
	colorful bool
}

func NewVibeLogger(out io.Writer, colorful bool) *VibeLogger {
	if out == nil {
		out = os.Stdout
	}
	return &VibeLogger{
		out:      out,
		colorful: colorful,
	}
}

func DefaultVibeLogger() *VibeLogger {
	return NewVibeLogger(os.Stdout, true)
}

func (l *VibeLogger) log(level EmotionalLevel, format string, args ...interface{}) {
	prefix := emotionalLevelText[level]

	if l.colorful {
		prefix = emotionalColor[level] + prefix + resetColor
	}

	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf(format, args...)

	fmt.Fprintf(l.out, "[%s] %s: %s\n", timestamp, prefix, message)

	if level == COOKED {
		os.Exit(1)
	}
}

func (l *VibeLogger) Fyi(format string, args ...interface{}) {
	l.log(FYI, format, args...)
}

func (l *VibeLogger) Uhoh(format string, args ...interface{}) {
	l.log(UHOH, format, args...)
}

func (l *VibeLogger) Crap(format string, args ...interface{}) {
	l.log(CRAP, format, args...)
}

func (l *VibeLogger) Cooked(format string, args ...interface{}) {
	l.log(COOKED, format, args...)
}

package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

// ANSI color codes
const (
	reset    = "\033[0m"
	bold     = "\033[1m"
	red      = "\033[31m"
	green    = "\033[32m"
	yellow   = "\033[33m"
	blue     = "\033[34m"
	magenta  = "\033[35m"
	cyan     = "\033[36m"
	gray     = "\033[37m"
	darkGray = "\033[90m"
	lightRed = "\033[91m"
)

// ColoredHandler is a slog.Handler that writes colored logs to an io.Writer.
type ColoredHandler struct {
	opts    slog.HandlerOptions
	w       io.Writer
	mu      sync.Mutex
	groups  []string
	attrs   []slog.Attr
	useJSON bool
}

// NewColoredHandler creates a new ColoredHandler that writes to w.
func NewColoredHandler(w io.Writer, opts *slog.HandlerOptions, useJSON bool) *ColoredHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &ColoredHandler{
		opts:    *opts,
		w:       w,
		useJSON: useJSON,
	}
}

// Enabled implements slog.Handler.
func (h *ColoredHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := h.opts.Level
	if minLevel == nil {
		return true
	}
	return level >= minLevel.Level()
}

// Handle implements slog.Handler.
//
//nolint:gocritic // Cannot change signature due to interface contract
func (h *ColoredHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	// If JSON format is requested, use the standard JSON handler with added color for level
	if h.useJSON {
		return h.handleJSON(&r)
	}

	// Format time
	timeStr := r.Time.Format("15:04:05.000")

	// Format level with color
	levelColor, levelStr := getLevelColor(r.Level)
	levelStr = fmt.Sprintf("%s%s%s", levelColor, levelStr, reset)

	// Start building the log line with time, level and message
	fmt.Fprintf(h.w, "%s %s %s\n", timeStr, levelStr, r.Message)

	// Format source if enabled - on a separate line
	if h.opts.AddSource && r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		if f.File != "" {
			// Extract just filename, not full path
			file := f.File
			line := f.Line
			fmt.Fprintf(h.w, "    %sSource: %s:%d%s\n", darkGray, file, line, reset)
		}
	}

	// Add attributes
	if len(h.attrs) > 0 || r.NumAttrs() > 0 {
		fmt.Fprintf(h.w, "    %sAttributes:%s\n", darkGray, reset)

		// Add handler attributes
		for _, attr := range h.attrs {
			fmt.Fprintf(h.w, "      %s%s%s: %v\n", cyan, attr.Key, reset, attr.Value.Any())
		}

		// Add record attributes
		r.Attrs(func(attr slog.Attr) bool {
			fmt.Fprintf(h.w, "      %s%s%s: %v\n", cyan, attr.Key, reset, attr.Value.Any())
			return true
		})
	}

	return nil
}

// handleJSON formats the log as JSON but with colored level
func (h *ColoredHandler) handleJSON(r *slog.Record) error {
	// Create a map for the JSON output
	m := make(map[string]any)

	// Add standard fields
	m["time"] = r.Time.Format(time.RFC3339)
	m["level"] = r.Level.String()
	m["msg"] = r.Message

	// Add source if enabled
	if h.opts.AddSource && r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		if f.File != "" {
			m["source"] = fmt.Sprintf("%s:%d", f.File, f.Line)
		}
	}

	// Add handler attributes
	for _, attr := range h.attrs {
		m[attr.Key] = attr.Value.Any()
	}

	// Add record attributes
	r.Attrs(func(attr slog.Attr) bool {
		m[attr.Key] = attr.Value.Any()
		return true
	})

	// Marshal to JSON
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	// Get color for level
	levelColor, _ := getLevelColor(r.Level)

	// Write colored JSON
	fmt.Fprintf(h.w, "%s%s%s\n", levelColor, string(b), reset)

	return nil
}

// WithAttrs implements slog.Handler.
func (h *ColoredHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Create a new handler without copying the mutex
	h2 := &ColoredHandler{
		opts:    h.opts,
		w:       h.w,
		groups:  append([]string{}, h.groups...),
		useJSON: h.useJSON,
	}

	// Copy and append attributes
	h2.attrs = make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(h2.attrs, h.attrs)
	copy(h2.attrs[len(h.attrs):], attrs)

	return h2
}

// WithGroup implements slog.Handler.
func (h *ColoredHandler) WithGroup(name string) slog.Handler {
	// Create a new handler without copying the mutex
	h2 := &ColoredHandler{
		opts:    h.opts,
		w:       h.w,
		attrs:   append([]slog.Attr{}, h.attrs...),
		useJSON: h.useJSON,
	}

	// Copy and append groups
	h2.groups = make([]string, len(h.groups)+1)
	copy(h2.groups, h.groups)
	h2.groups[len(h.groups)] = name

	return h2
}

// getLevelColor returns the color for the given level
func getLevelColor(level slog.Level) (colorCode, levelText string) {
	switch {
	case level >= slog.LevelError:
		return red, "ERROR"
	case level >= slog.LevelWarn:
		return yellow, "WARN "
	case level >= slog.LevelInfo:
		return green, "INFO "
	default:
		return blue, "DEBUG"
	}
}

package logger

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColoredHandler(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer

	// Create handler options
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Create a colored handler
	handler := NewColoredHandler(&buf, opts, false)
	logger := slog.New(handler)

	// Log messages at different levels
	logger.Debug("Debug message", "key", "value")
	logger.Info("Info message", "key", "value")
	logger.Warn("Warning message", "key", "value")
	logger.Error("Error message", "key", "value")

	// Get the output
	output := buf.String()

	// Verify that the output contains color codes
	assert.Contains(t, output, blue, "Output should contain blue color code")
	assert.Contains(t, output, green, "Output should contain green color code")
	assert.Contains(t, output, yellow, "Output should contain yellow color code")
	assert.Contains(t, output, red, "Output should contain red color code")

	// Verify that the output contains the messages
	assert.Contains(t, output, "Debug message", "Output should contain debug message")
	assert.Contains(t, output, "Info message", "Output should contain info message")
	assert.Contains(t, output, "Warning message", "Output should contain warning message")
	assert.Contains(t, output, "Error message", "Output should contain error message")

	// Verify that the output contains the attributes
	assert.Contains(t, output, "key", "Output should contain attribute key")
	assert.Contains(t, output, "value", "Output should contain attribute value")
}

func TestColoredHandlerJSON(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer

	// Create handler options
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Create a colored handler with JSON output
	handler := NewColoredHandler(&buf, opts, true)
	logger := slog.New(handler)

	// Log a message
	logger.Info("Test message", "key", "value")

	// Get the output
	output := buf.String()

	// Verify that the output contains color codes
	assert.Contains(t, output, green, "Output should contain green color code")

	// Verify that the output contains JSON fields
	assert.Contains(t, output, `"level": "INFO"`, "Output should contain level field")
	assert.Contains(t, output, `"msg": "Test message"`, "Output should contain msg field")
	assert.Contains(t, output, `"key": "value"`, "Output should contain key field")
}

func TestColoredHandlerWithAttrs(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer

	// Create handler options
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Create a colored handler
	baseHandler := NewColoredHandler(&buf, opts, false)

	// Add attributes with type assertion
	handlerWithAttrs := baseHandler.WithAttrs([]slog.Attr{
		slog.String("service", "test"),
		slog.Int("version", 1),
	})

	logger := slog.New(handlerWithAttrs)

	// Log a message
	logger.Info("Test message", "key", "value")

	// Get the output
	output := buf.String()

	// Verify that the output contains the attributes
	assert.Contains(t, output, "service", "Output should contain service attribute")
	assert.Contains(t, output, "test", "Output should contain test value")
	assert.Contains(t, output, "version", "Output should contain version attribute")
	assert.Contains(t, output, "1", "Output should contain version value")
	assert.Contains(t, output, "key", "Output should contain key attribute")
	assert.Contains(t, output, "value", "Output should contain value")
}

func TestColoredHandlerEnabled(t *testing.T) {
	// Create handler options
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	// Create a colored handler
	handler := NewColoredHandler(&bytes.Buffer{}, opts, false)

	// Check if levels are enabled
	assert.False(t, handler.Enabled(context.Background(), slog.LevelDebug), "Debug level should not be enabled")
	assert.True(t, handler.Enabled(context.Background(), slog.LevelInfo), "Info level should be enabled")
	assert.True(t, handler.Enabled(context.Background(), slog.LevelWarn), "Warn level should be enabled")
	assert.True(t, handler.Enabled(context.Background(), slog.LevelError), "Error level should be enabled")
}

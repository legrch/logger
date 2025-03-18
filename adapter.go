package logger

import (
	"fmt"
	"log/slog"
)

// LegacyLogger is the old logger interface, kept for backward compatibility
type LegacyLogger interface {
	// Debug logs a debug message
	Debug(args ...any)
	// Info logs an info message
	Info(args ...any)
	// Warn logs a warning message
	Warn(args ...any)
	// Error logs an error message
	Error(args ...any)
	// With adds structured context to the logger
	With(key string, value any) LegacyLogger
}

// LegacyAdapter adapts a *slog.Logger to the LegacyLogger interface
type LegacyAdapter struct {
	logger *slog.Logger
}

// NewLegacyAdapter creates a new LegacyAdapter
func NewLegacyAdapter(logger *slog.Logger) LegacyLogger {
	return &LegacyAdapter{logger: logger}
}

// Debug logs a debug message
func (l *LegacyAdapter) Debug(args ...any) {
	if len(args) == 0 {
		return
	}

	msg, attrs := extractMsgAndAttrs(args)
	l.logger.Debug(msg, attrs...)
}

// Info logs an info message
func (l *LegacyAdapter) Info(args ...any) {
	if len(args) == 0 {
		return
	}

	msg, attrs := extractMsgAndAttrs(args)
	l.logger.Info(msg, attrs...)
}

// Warn logs a warning message
func (l *LegacyAdapter) Warn(args ...any) {
	if len(args) == 0 {
		return
	}

	msg, attrs := extractMsgAndAttrs(args)
	l.logger.Warn(msg, attrs...)
}

// Error logs an error message
func (l *LegacyAdapter) Error(args ...any) {
	if len(args) == 0 {
		return
	}

	msg, attrs := extractMsgAndAttrs(args)
	l.logger.Error(msg, attrs...)
}

// With adds structured context to the logger
func (l *LegacyAdapter) With(key string, value any) LegacyLogger {
	return &LegacyAdapter{
		logger: l.logger.With(key, value),
	}
}

// extractMsgAndAttrs extracts the message and attributes from args
// The first argument is the message, and the rest are key-value pairs
func extractMsgAndAttrs(args []any) (msg string, attrs []any) {
	if len(args) == 0 {
		return "", nil
	}

	// First argument is the message
	msg, ok := args[0].(string)
	if !ok {
		msg = fmt.Sprint(args[0])
	}

	// Rest are key-value pairs
	attrs = make([]any, 0, len(args)-1)
	for i := 1; i < len(args); i++ {
		attrs = append(attrs, args[i])
	}

	return msg, attrs
}

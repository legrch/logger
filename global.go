package logger

import (
	"log/slog"
	"sync"
)

var (
	// defaultLogger is the default logger instance
	defaultLogger *slog.Logger

	// mu protects defaultLogger
	mu sync.RWMutex
)

// SetDefault sets the default logger instance
func SetDefault(logger *slog.Logger) {
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = logger
}

// Default returns the default logger instance
func Default() *slog.Logger {
	mu.RLock()
	defer mu.RUnlock()

	if defaultLogger == nil {
		// Return the standard library's default logger if the default logger is not set
		return slog.Default()
	}

	return defaultLogger
}

// Debug logs a debug message using the default logger
func Debug(msg string, args ...any) {
	Default().Debug(msg, args...)
}

// Info logs an info message using the default logger
func Info(msg string, args ...any) {
	Default().Info(msg, args...)
}

// Warn logs a warning message using the default logger
func Warn(msg string, args ...any) {
	Default().Warn(msg, args...)
}

// Error logs an error message using the default logger
func Error(msg string, args ...any) {
	Default().Error(msg, args...)
}

// With adds structured context to the default logger
func With(key string, value any) *slog.Logger {
	return Default().With(key, value)
}

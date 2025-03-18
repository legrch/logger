package logger

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

// MockLogger is a mock implementation of slog.Handler for testing
type MockLogger struct {
	mu     sync.Mutex
	logs   []LogEntry
	attrs  []slog.Attr
	groups []string
}

// LogEntry represents a log entry
type LogEntry struct {
	Level   slog.Level
	Message string
	Attrs   []slog.Attr
}

// NewMockLogger creates a new mock logger
func NewMockLogger() *slog.Logger {
	mock := &MockLogger{
		logs:  make([]LogEntry, 0),
		attrs: make([]slog.Attr, 0),
	}
	return slog.New(mock)
}

// Enabled implements slog.Handler.
func (*MockLogger) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

// Handle implements slog.Handler.
//
//nolint:gocritic // Cannot change signature due to interface contract
func (l *MockLogger) Handle(_ context.Context, r slog.Record) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Collect all attributes
	attrs := make([]slog.Attr, len(l.attrs))
	copy(attrs, l.attrs)

	r.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr)
		return true
	})

	// Add the log entry
	l.logs = append(l.logs, LogEntry{
		Level:   r.Level,
		Message: r.Message,
		Attrs:   attrs,
	})

	return nil
}

// WithAttrs implements slog.Handler.
func (l *MockLogger) WithAttrs(attrs []slog.Attr) slog.Handler {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Create a new logger that shares the logs slice
	newLogger := &MockLogger{
		logs:   l.logs,
		attrs:  make([]slog.Attr, len(l.attrs)+len(attrs)),
		groups: make([]string, len(l.groups)),
	}

	copy(newLogger.attrs, l.attrs)
	copy(newLogger.attrs[len(l.attrs):], attrs)
	copy(newLogger.groups, l.groups)

	return newLogger
}

// WithGroup implements slog.Handler.
func (l *MockLogger) WithGroup(name string) slog.Handler {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Create a new logger that shares the logs slice
	newLogger := &MockLogger{
		logs:   l.logs,
		attrs:  make([]slog.Attr, len(l.attrs)),
		groups: make([]string, len(l.groups)+1),
	}

	copy(newLogger.attrs, l.attrs)
	copy(newLogger.groups, l.groups)
	newLogger.groups[len(l.groups)] = name

	return newLogger
}

// GetLogs returns all log entries
func (l *MockLogger) GetLogs() []LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()

	logs := make([]LogEntry, len(l.logs))
	copy(logs, l.logs)

	return logs
}

// GetLogsByLevel returns log entries filtered by level
func (l *MockLogger) GetLogsByLevel(level slog.Level) []LogEntry {
	l.mu.Lock()
	defer l.mu.Unlock()

	logs := make([]LogEntry, 0)
	for _, log := range l.logs {
		if log.Level == level {
			logs = append(logs, log)
		}
	}

	return logs
}

// Clear clears all log entries
func (l *MockLogger) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.logs = make([]LogEntry, 0)
}

// String returns a string representation of the log entries
func (l *MockLogger) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	result := ""
	for _, log := range l.logs {
		result += fmt.Sprintf("[%s] %s %v\n", log.Level, log.Message, log.Attrs)
	}

	return result
}

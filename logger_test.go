package logger

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	// Test with default config
	cfg := &Config{
		Level:            "info",
		Format:           "json",
		EnableCaller:     true,
		EnableStacktrace: true,
	}

	log, err := New(cfg)
	require.NoError(t, err)
	assert.NotNil(t, log)
	assert.IsType(t, &slog.Logger{}, log)

	// Test with invalid level
	cfg.Level = "invalid"
	log, err = New(cfg)
	require.Error(t, err)
	assert.Nil(t, log)
}

func TestGlobalLogger(t *testing.T) {
	// Test default logger (should be slog.Default())
	defaultLog := Default()
	assert.NotNil(t, defaultLog)

	// Set a mock logger
	mock := NewMockLogger()
	SetDefault(mock)

	// Test global functions
	Debug("debug message")
	Info("info message")
	Warn("warn message")
	Error("error message")

	// Get logs from mock
	mockHandler := mock.Handler().(*MockLogger)
	logs := mockHandler.GetLogs()
	assert.Len(t, logs, 4)
	assert.Equal(t, slog.LevelDebug, logs[0].Level)
	assert.Equal(t, slog.LevelInfo, logs[1].Level)
	assert.Equal(t, slog.LevelWarn, logs[2].Level)
	assert.Equal(t, slog.LevelError, logs[3].Level)
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		level    string
		expected slog.Level
		hasError bool
	}{
		{"debug", slog.LevelDebug, false},
		{"info", slog.LevelInfo, false},
		{"warn", slog.LevelWarn, false},
		{"error", slog.LevelError, false},
		{"invalid", slog.LevelInfo, true},
	}

	for _, test := range tests {
		t.Run(test.level, func(t *testing.T) {
			level, err := parseLogLevel(test.level)
			if test.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, level)
			}
		})
	}
}

func TestExtractMsgAndAttrs(t *testing.T) {
	tests := []struct {
		name      string
		args      []any
		wantMsg   string
		wantAttrs []any
	}{
		{
			name:      "empty args",
			args:      []any{},
			wantMsg:   "",
			wantAttrs: nil,
		},
		{
			name:      "message only",
			args:      []any{"message"},
			wantMsg:   "message",
			wantAttrs: []any{},
		},
		{
			name:      "message with one key-value pair",
			args:      []any{"message", "key", "value"},
			wantMsg:   "message",
			wantAttrs: []any{"key", "value"},
		},
		{
			name:      "message with multiple key-value pairs",
			args:      []any{"message", "key1", "value1", "key2", 42},
			wantMsg:   "message",
			wantAttrs: []any{"key1", "value1", "key2", 42},
		},
		{
			name:      "non-string message",
			args:      []any{42, "key", "value"},
			wantMsg:   "42",
			wantAttrs: []any{"key", "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsg, gotAttrs := extractMsgAndAttrs(tt.args)
			assert.Equal(t, tt.wantMsg, gotMsg)
			assert.Equal(t, tt.wantAttrs, gotAttrs)
		})
	}
}

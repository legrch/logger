package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMsgAndAttrsFunction(t *testing.T) {
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

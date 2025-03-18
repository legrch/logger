package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Config holds all the logger-related configuration
type Config struct {
	// Level is the minimum enabled logging level
	Level string `envconfig:"LEVEL" default:"info"`

	// Format is the log format ("json" or "console")
	Format string `envconfig:"FORMAT" default:"json"`

	// EnableCaller adds the file:line caller info to log output
	EnableCaller bool `envconfig:"ENABLE_CALLER" default:"true"`

	// EnableStacktrace enables automatic stacktrace capturing
	EnableStacktrace bool `envconfig:"ENABLE_STACKTRACE" default:"true"`

	// EnableColors enables colored output for console format
	EnableColors bool `envconfig:"ENABLE_COLORS" default:"false"`

	// Environment is the current environment (development, staging, production)
	Environment string `envconfig:"ENVIRONMENT" default:"production"`
}

// New creates a new slog.Logger with the given configuration
func New(cfg *Config) (*slog.Logger, error) {
	// Parse log level
	level, err := parseLogLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	// Create handler options
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.EnableCaller,
	}

	// Create handler based on format
	var handler slog.Handler

	// Check if we should use colored output
	useColors := cfg.EnableColors || isLocalEnvironment(cfg.Environment)

	if cfg.Format == "console" && useColors {
		// Use colored handler for console format in local environment
		handler = NewColoredHandler(os.Stdout, opts, false)
	} else if cfg.Format == "json" && useColors {
		// Use colored JSON handler
		handler = NewColoredHandler(os.Stdout, opts, true)
	} else if cfg.Format == "console" {
		// Use standard text handler
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		// Use standard JSON handler
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	// Create logger
	return slog.New(handler), nil
}

// isLocalEnvironment checks if the environment is a local/development environment
func isLocalEnvironment(env string) bool {
	env = strings.ToLower(env)
	return env == "development" || env == "local" || env == "dev"
}

// parseLogLevel parses a log level string into a slog.Level
func parseLogLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unknown log level: %s", level)
	}
}

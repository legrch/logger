package logger

import (
	"fmt"
)

// Format constants
const (
	FormatJSON    = "json"
	FormatConsole = "console"
)

// Init initializes the logger with the given configuration and sets it as the default logger
func Init(cfg *Config) error {
	// Create a new logger
	log, err := New(cfg)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}

	// Set the default logger
	SetDefault(log)

	// Log initialization
	log.Info("Logger initialized",
		"level", cfg.Level,
		"format", cfg.Format,
		"enableCaller", cfg.EnableCaller,
		"enableStacktrace", cfg.EnableStacktrace,
		"environment", cfg.Environment,
		"enableColors", cfg.EnableColors,
	)

	return nil
}

// InitFromEnv initializes the logger from environment variables
func InitFromEnv(env string) error {
	// Set format to console in development environment
	format := FormatJSON
	enableColors := false

	if env == "development" || env == "local" || env == "dev" {
		format = FormatConsole
		enableColors = true
	}

	// Create config
	cfg := &Config{
		Level:            "info",
		Format:           format,
		EnableCaller:     true,
		EnableStacktrace: true,
		EnableColors:     enableColors,
		Environment:      env,
	}

	// Initialize logger
	return Init(cfg)
}

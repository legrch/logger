package main

import (
	"fmt"
	"os"

	"github.com/legrch/logger/pkg/logger"
)

func main() {
	// Create a logger configuration with colored output
	cfg := &logger.Config{
		Level:            "debug",
		Format:           "console",
		EnableCaller:     true,
		EnableStacktrace: true,
		EnableColors:     true,
		Environment:      "development",
	}

	// Initialize the logger
	err := logger.Init(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Log messages at different levels
	logger.Debug("This is a debug message", "count", 1)
	logger.Info("This is an info message", "user", "john")
	logger.Warn("This is a warning message", "latency", "100ms")
	logger.Error("This is an error message", "error", "connection refused")

	// Log with structured context
	userLogger := logger.With("user_id", "123")
	userLogger.Info("User logged in", "ip", "192.168.1.1")

	// Log with nested attributes
	orderLogger := userLogger.With("order_id", "ABC123")
	orderLogger.Info("Order processed", "amount", 99.99, "currency", "USD")

	fmt.Println("\nJSON format example:")

	// Create a logger with JSON format but still colored
	jsonCfg := &logger.Config{
		Level:        "debug",
		Format:       "json",
		EnableColors: true,
		Environment:  "development",
	}

	jsonLogger, err := logger.New(jsonCfg)
	if err != nil {
		fmt.Printf("Failed to create JSON logger: %v\n", err)
		os.Exit(1)
	}

	// Log with JSON format
	jsonLogger.Info("This is a JSON log", "key", "value")
}

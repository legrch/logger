package main

import (
	"fmt"
	"os"
	"time"

	"github.com/legrch/logger/pkg/logger"
)

func main() {
	fmt.Println("=== With Source Information ===")

	// Create a logger with source information
	withSourceCfg := &logger.Config{
		Level:        "debug",
		Format:       "console",
		EnableCaller: true,
		EnableColors: true,
		Environment:  "development",
	}

	// Initialize the logger
	withSourceLogger, err := logger.New(withSourceCfg)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Log messages at different levels
	withSourceLogger.Debug("Debug message with source info", "key", "value")
	withSourceLogger.Info("Info message with source info", "key", "value")
	withSourceLogger.Warn("Warning message with source info", "key", "value")
	withSourceLogger.Error("Error message with source info", "key", "value")

	fmt.Println("\n=== Without Source Information ===")

	// Create a logger without source information
	withoutSourceCfg := &logger.Config{
		Level:        "debug",
		Format:       "console",
		EnableCaller: false,
		EnableColors: true,
		Environment:  "development",
	}

	// Initialize the logger
	withoutSourceLogger, err := logger.New(withoutSourceCfg)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Log messages at different levels
	withoutSourceLogger.Debug("Debug message without source info", "key", "value")
	withoutSourceLogger.Info("Info message without source info", "key", "value")
	withoutSourceLogger.Warn("Warning message without source info", "key", "value")
	withoutSourceLogger.Error("Error message without source info", "key", "value")

	fmt.Println("\n=== JSON vs Text Format ===")

	// Create a JSON-formatted logger
	jsonCfg := &logger.Config{
		Level:  "info",
		Format: "json",
	}

	jsonLogger, err := logger.New(jsonCfg)
	if err != nil {
		fmt.Printf("Failed to create JSON logger: %v\n", err)
		os.Exit(1)
	}

	jsonLogger.Info("json logger example",
		"user_id", 12345,
		"request_id", "abc-123-xyz",
		"timestamp", time.Now().Format(time.RFC3339),
	)

	// Add a newline for readability
	os.Stdout.Write([]byte("\n"))

	// Create a text-formatted logger
	textCfg := &logger.Config{
		Level:  "info",
		Format: "text",
	}

	textLogger, err := logger.New(textCfg)
	if err != nil {
		fmt.Printf("Failed to create text logger: %v\n", err)
		os.Exit(1)
	}

	textLogger.Info("text logger example",
		"user_id", 12345,
		"request_id", "abc-123-xyz",
		"timestamp", time.Now().Format(time.RFC3339),
	)
}

# Go Structured Logger

[![Go Reference](https://pkg.go.dev/badge/github.com/legrch/logger.svg)](https://pkg.go.dev/github.com/legrch/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/legrch/logger)](https://goreportcard.com/report/github.com/legrch/logger)
[![License](https://img.shields.io/github/license/legrch/logger)](LICENSE)
[![Release](https://img.shields.io/github/v/release/legrch/logger)](https://github.com/legrch/logger/releases)

A structured logging package for Go applications built on top of the standard library's [slog](https://pkg.go.dev/log/slog) package.

## Features

- **Built on Go's standard library**: Uses `slog` from Go 1.21+
- **Structured logging**: Key-value pairs for better log filtering and analysis
- **Multiple output formats**: JSON for production, text for development
- **Colored console output**: Improves readability during local development
- **Customizable log levels**: debug, info, warn, error
- **Source information**: Automatically includes caller information (file:line)
- **Global logger**: Convenient access throughout your application
- **Testing support**: Mock logger for easy testing

## Installation

```bash
go get github.com/legrch/logger
```

## Quick Start

```go
package main

import (
	"github.com/legrch/logger"
	"os"
)

func main() {
	// Initialize logger with default configuration
	cfg := &logger.Config{
		Level:           "info",
		Format:          "text",
		Color:           true,
		CallerEnabled:   true,
		CallerSkipFrame: 1,
	}
	
	log := logger.New(cfg, os.Stdout)
	
	// Use the logger
	log.Info("application started", "version", "1.0.0")
	log.Debug("debug message") // Won't be shown with info level
	log.Error("something went wrong", "error", "connection refused", "retry", true)
	
	// With structured fields
	log.With("request_id", "abc123").Info("processing request")
	
	// Set as global logger
	logger.SetGlobal(log)
	
	// Use global logger elsewhere
	globalLogger := logger.Global()
	globalLogger.Info("using global logger")
}
```

## Documentation

For detailed documentation, examples, and API reference, please visit:

- [Package Documentation](https://pkg.go.dev/github.com/legrch/logger)
- [Examples](https://github.com/legrch/logger/tree/main/pkg/logger/examples)
- [Logger Package Documentation](pkg/logger/README.md)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 
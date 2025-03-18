# Logger Package

This package provides a structured logging solution for the application using [slog](https://pkg.go.dev/log/slog) from the Go standard library.

## Features

- Direct use of `*slog.Logger` from the standard library
- Structured logging with key-value pairs
- Multiple log levels (debug, info, warn, error)
- JSON or text output formats
- Colored console output for local development
- Caller information (file:line)
- Global logger instance
- Mock logger for testing
- Legacy adapter for backward compatibility

## Usage

### Basic Usage

```go
package main

import (
    "github.com/legrch/logger"
)

func main() {
    // Initialize logger with default configuration
    cfg := &logger.Config{
        Level:           "info",
        Format:          "console", // Use "json" for production
        EnableCaller:    true,
        EnableStacktrace: true,
    }
    
    if err := logger.Init(cfg); err != nil {
        panic(err)
    }
    
    // Use global logger functions
    logger.Info("Application started")
    logger.Debug("Debug message") // Won't be logged if level is info
    logger.Warn("Warning message", "code", 123)
    logger.Error("Error message", "error", err)
    
    // Add structured context
    userLogger := logger.With("user_id", "123")
    userLogger.Info("User logged in")
    
    // Get the underlying *slog.Logger
    slogLogger := logger.Default()
    
    // Use slog.Logger directly
    slogLogger.Info("Direct slog usage", "key", "value")
}
```

### Configuration from Environment

The logger can be configured from environment variables:

```go
// In config.go
type Config struct {
    // ...
    Logger logger.Config `envconfig:"LOGGER" required:"true"`
    // ...
}

// In main.go
cfg, err := config.LoadConfig()
if err != nil {
    panic(err)
}

// Logger is automatically initialized in LoadConfig
```

Environment variables for logger configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `LOGGER_LEVEL` | Minimum log level (debug, info, warn, error) | info |
| `LOGGER_FORMAT` | Log format (json, console) | json |
| `LOGGER_ENABLE_CALLER` | Include caller information | true |
| `LOGGER_ENABLE_STACKTRACE` | Include stack traces for errors | true |
| `LOGGER_ENABLE_COLORS` | Enable colored output for console format | false |
| `LOGGER_ENVIRONMENT` | Current environment (development, staging, production) | production |

### Colored Logging

In development environments, the logger automatically uses colored output for better readability:

- **Blue** for debug messages
- **Green** for info messages
- **Yellow** for warning messages
- **Red** for error messages

The colored output format is designed for maximum readability:
- Timestamp and log level are color-coded
- Message appears on the same line as the level for quick scanning
- Source information (file:line) is displayed in gray on a separate line
- Attributes are indented and displayed with colored keys

Example output:
```
23:38:52.913 INFO  This is a log message
    Source: /path/to/file.go:42
    Attributes:
      user_id: 123
      action: login
```

You can enable colored logging in any environment by setting:
```
LOGGER_ENABLE_COLORS=true
```

Or by setting the environment to development:
```
LOGGER_ENVIRONMENT=development
```

### Testing with Mock Logger

```go
package mypackage_test

import (
    "log/slog"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/legrch/logger"
)

func TestMyFunction(t *testing.T) {
    // Create a mock logger
    mock := logger.NewMockLogger()
    
    // Set it as the default logger
    logger.SetDefault(mock)
    
    // Call function that uses logger
    MyFunction()
    
    // Get the mock handler
    mockHandler := mock.Handler().(*logger.MockLogger)
    
    // Assert logs
    logs := mockHandler.GetLogs()
    assert.Len(t, logs, 1)
    assert.Equal(t, slog.LevelInfo, logs[0].Level)
    assert.Equal(t, "Function called", logs[0].Message)
}
```

### Legacy Adapter for Backward Compatibility

If you have existing code that uses the old logger interface, you can use the legacy adapter:

```go
// Create a legacy adapter
legacyLogger := logger.NewLegacyAdapter(logger.Default())

// Use the legacy interface
legacyLogger.Info("This uses the old interface", "key", "value")
```

## Log Levels

- **Debug**: Detailed information, typically of interest only when diagnosing problems
- **Info**: Confirmation that things are working as expected
- **Warn**: Indication that something unexpected happened, or may happen in the near future
- **Error**: Due to a more serious problem, the software has not been able to perform some function 
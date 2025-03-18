# Logger Examples

This directory contains examples demonstrating how to use the `github.com/legrch/logger` package in different scenarios.

## Overview

The examples showcase various features of the logger package, including colored output, different formatting options, and structured logging patterns.

## Usage

### Prerequisites
- Go 1.21 or later
- The `github.com/legrch/logger` package

### Examples

#### Colored Logs

The `colored` example demonstrates how to use colored logging output for improved readability in terminal environments.

```bash
# Run the colored logs example
cd colored
go run main.go
```

Key features demonstrated:
- Colored output for different log levels
- Structured logging with key-value pairs
- Nested context with `.With()`
- JSON formatting with colors

#### Formatting Options

The `formatting` example shows different formatting options available in the logger.

```bash
# Run the formatting example
cd formatting
go run main.go
```

Key features demonstrated:
- Source code information in logs
- Different logging formats (console vs. JSON)
- Debug, info, warning, and error level messages
- Setting different configuration options

## Best Practices

When using the logger package, consider these best practices:

1. Use structured logging (key-value pairs) instead of string interpolation
2. Add context to logs using the `.With()` method
3. Use appropriate log levels based on the importance of the message
4. Configure the logger based on your environment (colored text for development, JSON for production)
5. Include relevant contextual information in logs (request IDs, user IDs, etc.)

## Related Documentation

- [Logger Package Documentation](https://pkg.go.dev/github.com/legrch/logger/pkg/logger)
- [Go's slog Package](https://pkg.go.dev/log/slog) - The standard library package this logger builds upon 
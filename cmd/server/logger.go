package main

import (
	"log/slog"
	"os"
)

func initLogger(level string) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLogLevel(level),
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}

func parseLogLevel(level string) slog.Leveler {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

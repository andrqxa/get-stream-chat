package main

import (
	"get-stream-chat/internal/app"
	"get-stream-chat/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("config reading error")
	}

	initLogger(cfg.LogLevel)

	slog.Info("application starting")

	app, err := app.New(cfg)
	if err != nil {
		slog.Error("failed to init app", "error", err)
		os.Exit(1)
	}

	if err := app.Run(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/garrettladley/garrettladley/internal/is/server"
	"github.com/garrettladley/garrettladley/internal/is/settings"
	"github.com/garrettladley/garrettladley/pkg/ai/openai"
	pserver "github.com/garrettladley/garrettladley/pkg/server"
	"github.com/garrettladley/garrettladley/pkg/xslog"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())

	settings, err := settings.Load()
	if err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"failed to load settings",
			xslog.Error(err),
		)
		os.Exit(1)
	}

	app := server.New(&server.Config{
		Client: openai.New(settings.AI.Key),
		Config: pserver.Config{
			Logger: logger,
		},
	})

	go func() {
		if err := app.Listen(":8080"); err != nil {
			slog.LogAttrs(
				ctx,
				slog.LevelError,
				"failed to start server",
				xslog.Error(err),
			)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	slog.LogAttrs(
		ctx,
		slog.LevelInfo,
		"stopping server",
	)
	cancel()

	if err := app.Shutdown(); err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"failed to shutdown server",
			xslog.Error(err),
		)
	}

	slog.LogAttrs(
		ctx,
		slog.LevelInfo,
		"server shutdown",
	)
}

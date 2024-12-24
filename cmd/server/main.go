package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/garrettladley/garrettladley/internal/api/is/conf"
	sserver "github.com/garrettladley/garrettladley/internal/site/server"
	"github.com/garrettladley/garrettladley/pkg/server"
	"github.com/garrettladley/garrettladley/pkg/xslog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())

	_, err := conf.Load()
	if err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"failed to load settings",
			xslog.Error(err),
		)
		os.Exit(1)
	}

	app := sserver.New(&sserver.Config{
		Config: server.Config{
			Logger: logger,
		},
		StaticFn: static,
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

func static(app *fiber.App) {
	app.Get("/public/*", adaptor.HTTPHandler(public()))
}

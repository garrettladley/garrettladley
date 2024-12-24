package server

import (
	"github.com/garrettladley/garrettladley/pkg/server"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Config   server.Config
	StaticFn func(*fiber.App)
}

func New(cfg *Config) *fiber.App {
	app := server.New(cfg.Config, func(app *fiber.App) {
		registerRoutes(app)
		cfg.StaticFn(app)
	})
	return app
}

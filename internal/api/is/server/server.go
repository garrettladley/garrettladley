package server

import (
	"github.com/garrettladley/garrettladley/pkg/ai"
	"github.com/garrettladley/garrettladley/pkg/server"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Client ai.AI
	Config server.Config
}

func New(cfg *Config) *fiber.App {
	app := server.New(cfg.Config, func(app *fiber.App) { registerRoutes(app, cfg) })
	return app
}

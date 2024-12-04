package server

import (
	"github.com/garrettladley/garrettladley/pkg/ai/openai"
	"github.com/garrettladley/garrettladley/pkg/server"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Client *openai.Client
	Config server.Config
}

func New(cfg *Config) *fiber.App {
	app := server.New(cfg.Config)
	registerRoutes(app, cfg)

	return app
}
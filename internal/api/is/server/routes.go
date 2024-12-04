package server

import (
	"github.com/garrettladley/garrettladley/internal/api/is/handlers/even"
	"github.com/garrettladley/garrettladley/internal/api/is/handlers/odd"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App, cfg *Config) {
	apiIs := app.Group("/api/is")

	even.Routes(apiIs, cfg.Client)
	odd.Routes(apiIs, cfg.Client)
}

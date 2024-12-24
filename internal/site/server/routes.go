package server

import (
	"github.com/garrettladley/garrettladley/internal/site/handlers"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	handlers.Routes(app)
}

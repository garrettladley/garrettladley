package handlers

import (
	"github.com/garrettladley/garrettladley/internal/site/views/home"
	"github.com/garrettladley/garrettladley/pkg/xtempl"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Home(c *fiber.Ctx) error {
	return xtempl.Render(c, home.Index())
}

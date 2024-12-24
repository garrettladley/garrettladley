package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router) {
	s := New()

	r.Get("/", s.Home)
}

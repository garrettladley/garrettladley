package odd

import (
	"github.com/garrettladley/garrettladley/pkg/ai"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, ai ai.AI) {
	s := New(ai)

	r.Route("/odd", func(r fiber.Router) {
		r.Get("/:n", s.Is)
	})
}

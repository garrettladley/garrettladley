package odd

import (
	"github.com/garrettladley/garrettladley/pkg/ai/openai"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, client *openai.Client) {
	s := newService(client)

	r.Route("/odd", func(r fiber.Router) {
		r.Get("/:n", s.Is)
	})
}

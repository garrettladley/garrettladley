package odd

import (
	"strconv"

	"github.com/garrettladley/garrettladley/internal/api/is/predicate"
	"github.com/garrettladley/garrettladley/pkg/xerr"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Is(c *fiber.Ctx) error {
	n, err := strconv.ParseInt(c.Params("n"), 10, 64)
	if err != nil {
		return xerr.BadRequest(err)
	}
	var result bool
	if c.QueryBool("ai", false) {
		result, err = s.client.Is(c.Context(), n, predicate.Odd)
		if err != nil {
			return err
		}
	} else {
		result = n%2 != 0
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"is_odd": result})
}

package server

import (
	"log/slog"
	"net/http"
	"time"

	go_json "github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"

	_ "embed"

	"github.com/garrettladley/garrettladley/pkg/xerr"
)

type Config struct {
	Logger *slog.Logger
}

func New(cfg Config, register func(*fiber.App)) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:       go_json.Marshal,
		JSONDecoder:       go_json.Unmarshal,
		ErrorHandler:      xerr.ErrorHandler,
		PassLocalsToViews: true,
	})
	setupMiddleware(app, cfg)
	setupHealthCheck(app)
	setupRobotsTxt(app)
	setupSiteMap(app)
	register(app)
	setup404Handler(app)

	return app
}

func setupMiddleware(app *fiber.App, cfg Config) {
	app.Use(recover.New())
	app.Use(slogfiber.New(cfg.Logger))
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

func setupHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
}

//go:embed artifacts/robots.txt
var robotsTxt string

func setupRobotsTxt(app *fiber.App) {
	app.Get("/robots.txt", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(robotsTxt)
	})
}

//go:embed artifacts/sitemap.xml
var sitemapXml string

func setupSiteMap(app *fiber.App) {
	app.Get("/sitemap.xml", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(sitemapXml)
	})
}

func setup404Handler(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Page not found",
			"path":  c.Path(),
		})
	})
}

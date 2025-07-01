package router

import (
	"time"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/smilepakawat/gofiber-observability/internal/handler"
	"github.com/smilepakawat/gofiber-observability/internal/middleware"
)

func Route(app *fiber.App) {
	app.Use(otelfiber.Middleware())

	health := app.Group("/health")
	health.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now(),
		})
	})

	api := app.Group("/api")
	api.Use(middleware.Logger())
	api.Get("hello", handler.HelloWorld)
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smilepakawat/gofiber-observability/internal/handler"
	"github.com/smilepakawat/gofiber-observability/internal/middleware"
)

func Route(app *fiber.App) {
	api := app.Group("/api")
	api.Use(middleware.Logger())
	api.Get("hello", handler.HelloWorld)
}

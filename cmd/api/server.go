package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smilepakawat/gofiber-observability/internal/router"
)

func StartServer() {
	app := fiber.New()

	router.Route(app)

	app.Listen(":3000")
}

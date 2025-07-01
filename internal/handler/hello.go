package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

func HelloWorld(c *fiber.Ctx) error {
	tracer := otel.GetTracerProvider().Tracer("hello-handler")

	_, span := tracer.Start(c.UserContext(), "handler")
	defer span.End()

	time.Sleep(100 * time.Millisecond)

	return c.SendString("Hello World!")
}

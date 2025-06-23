package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func HelloWorld(c *fiber.Ctx) error {
	tracer := otel.Tracer("main-handler")

	_, span := tracer.Start(c.UserContext(), "perform-work")
	defer span.End()

	time.Sleep(100 * time.Millisecond)
	span.AddEvent("Work finished!")

	span.SetAttributes(attribute.String("my.custom.value", "hello world"))

	return c.SendString("Hello World!")
}

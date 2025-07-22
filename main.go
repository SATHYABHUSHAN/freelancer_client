package main

import (
	_ "github.com/SATHYABHUSHAN/freelancer_client/docs" // Swagger docs
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Freelancer Platform API
// @version 1.0
// @description APIs for Freelancers & Clients
// @host localhost:8080
// @BasePath /

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber backend!")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // Swagger endpoint

	app.Listen(":8080")
}

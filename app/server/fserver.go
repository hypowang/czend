package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	app.Use(logger.New())

	app.Static("/", "/root/go/src/czend/web/public")

	app.Get("/hello", hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":80"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

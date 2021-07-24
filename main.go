package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello V2.4, World 👋!")
	})

	app.Listen(":3000")
}

func Add(a int, b int) int {
	return a + b
}

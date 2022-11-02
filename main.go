package main

import (
	"bootmind/pkg/connector"
	"bootmind/pkg/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New(fiber.Config{
        BodyLimit: 1024 * 1024 * 100, // 100MB
    })

	db := connector.Connect()

            // contains index.html outside the app folder
	app.Static("/public", "./public")

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(map[string]string{
            "message": "Figo API",
        })
    })


	route.Expenses(app, db)
    app.Listen(":3000")
}
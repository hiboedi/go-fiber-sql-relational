package main

import (
	"go-fiber-relationaldb/database"
	"go-fiber-relationaldb/database/migrations"
	"go-fiber-relationaldb/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// connect db
	database.DatabaseConnect()
	// auto Migrate db
	migrations.Migration()
	// fiber init
	app := fiber.New()
	// router init
	routes.RouteInit(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Ok",
			"age":     23,
		})
	})

	app.Listen(":8000")
}

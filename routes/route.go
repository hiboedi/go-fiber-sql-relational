package routes

import (
	"go-fiber-relationaldb/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/users", controllers.UserGetAll)

	app.Post("/users", controllers.CreateUser)

	app.Get("/lockers", controllers.LockerGetAll)

	app.Post("/lockers", controllers.CreateLocker)

	app.Get("/posts", controllers.PostGetAll)

	app.Post("/posts", controllers.CreatePost)

	app.Get("/tags", controllers.TagGetAll)

	app.Post("/tags", controllers.CreateTag)

}

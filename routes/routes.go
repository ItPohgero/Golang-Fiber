package routes

import (
	"Golang-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Welcome)
	route := app.Group(prefix)
	{
		route.Post(register, controllers.Register)
		route.Post(login, controllers.Login)

		// with header authorization bearer token
		route.Use(controllers.IsAuthorized)
		{
			route.Get(users, controllers.Users)
		}
	}
}

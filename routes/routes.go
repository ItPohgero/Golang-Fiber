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

		route.Use(controllers.IsAuthorized)
		{
			route.Get(usersList, controllers.UsersList)
			route.Get(userShow, controllers.UserShow)
			route.Patch(userUpdate, controllers.UserUpdate)
			route.Delete(userDestroy, controllers.UserDestroy)

			route.Get(blogShow, controllers.BlogShow)

		}
	}
}

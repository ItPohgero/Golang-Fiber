package routes

import (
	"Golang-fiber/controllers"
	"Golang-fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Welcome)
	route := app.Group(prefix)
	{
		route.Post(register, controllers.Register)
		route.Post(login, controllers.Login)

		route.Get(usersList, middleware.Protected(), controllers.UsersList)
		route.Get(userShow, middleware.Protected(), controllers.UserShow)
		route.Patch(userUpdate, middleware.Protected(), controllers.UserUpdate)
		route.Delete(userDestroy, middleware.Protected(), controllers.UserDestroy)

		route.Get(blogShow, middleware.Protected(), controllers.BlogShow)
	}
}

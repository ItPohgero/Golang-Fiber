package routes

import (
	"Golang-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Welcome)
	auth := app.Group("/api/v1")
	{
		auth.Post("/register", controllers.Register)
		auth.Post("/login", controllers.Login)
	}

	//middleware for authentication and authorization for all routes header Authorization: Bearer <token>
	auth.Use(controllers.IsAuthorized)
	{
		user := app.Group("/api/v1")
		{
			user.Get("/users", controllers.UserList)
		}
	}

}

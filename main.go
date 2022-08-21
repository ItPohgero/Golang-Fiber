package main

import (
	"Golang-fiber/database"
	"Golang-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Authorization",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := app.Listen(":" + port)
	if err != nil {
		return
	}
}

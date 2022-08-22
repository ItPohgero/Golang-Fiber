package controllers

import (
	"Golang-fiber/database"
	"Golang-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func BlogShow(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog models.Blog
	database.DB.First(&blog, id)

	var user models.User
	database.DB.First(&user, blog.UserID)
	return c.JSON(fiber.Map{
		"user": user,
		"blog": blog,
	})
}
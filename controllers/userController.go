package controllers

import (
	"Golang-fiber/database"
	"Golang-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func Users(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	return c.JSON(users)
}

func User(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)

	return c.JSON(user)
}

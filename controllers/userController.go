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

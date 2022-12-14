package controllers

import (
	"Golang-fiber/database"
	"Golang-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func UsersList(c *fiber.Ctx) error {

	var users []models.User
	//database.DB.Preload(clause.Associations).Find(&users)
	database.DB.Preload("Blogs").Find(&users)
	return c.JSON(fiber.Map{
		"message": "Users found",
		"data":    users,
	})
}

func UserShow(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.Preload("Blogs").First(&user, id)

	return c.JSON(fiber.Map{
		"message": "User found",
		"data":    user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(
			fiber.Map{
				"message": "User not found",
			})
	}

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"message": "Error parsing body",
			})
	}

	database.DB.Model(&user).Updates(
		models.User{
			Name:  data["name"],
			Email: data["email"],
		})
	return c.JSON(fiber.Map{
		"message": "User updated",
		"user":    user,
	})
}

func UserDestroy(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(
			fiber.Map{
				"message": "User not found",
			})
	}

	database.DB.Delete(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted",
		"user":    user,
	})
}

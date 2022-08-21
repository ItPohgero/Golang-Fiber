package controllers

import (
	"Golang-fiber/database"
	"Golang-fiber/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	return c.JSON(users)
}

func IsAuthorized(c *fiber.Ctx) error {

	token := c.Get("Authorization")
	if token == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(
			fiber.Map{
				"message": "Unauthorized",
			})
	}

	tokenString := token[7:]
	tokenClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(
			fiber.Map{
				"message": "Unauthorized",
			})
	}

	return c.Next()
}

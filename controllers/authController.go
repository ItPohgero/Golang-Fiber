package controllers

import (
	"Golang-fiber/database"
	"Golang-fiber/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const ScretKey = "secret"

func Register(c *fiber.Ctx) error {
	// Parse body
	var data map[string]string

	// Validate data is correct or not null and length is greater than 0
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"message": "Error parsing body",
			})
	}

	// create bcrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	// save user to database
	database.DB.Create(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func Login(c *fiber.Ctx) error {
	// Parse body
	var data map[string]string

	// Validate data is correct or not null and length is greater than 0
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(
			fiber.Map{
				"message": "Error parsing body",
			})
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(
			fiber.Map{
				"message": "User not found",
			})
	}
	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(
			fiber.Map{
				"message": "Invalid password",
			})
	}

	claims := jwt.MapClaims{
		"id":   strconv.Itoa(int(user.Id)),
		"user": user.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ScretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(
			fiber.Map{
				"message": "Error creating token",
			})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   tokenString,
	})
}

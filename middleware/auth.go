package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v2"
)

func Protected() func(*fiber.Ctx) error {
	return jwt.New(jwt.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized | Missing or malformed Token",
		})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized | Invalid or expired Token",
		})
	}
}

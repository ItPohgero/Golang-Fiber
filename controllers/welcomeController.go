package controllers

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendString("Hello, World 👋!")
}

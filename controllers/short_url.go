package controllers

import "github.com/gofiber/fiber/v2"

func ShortURL(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
        "message": "Shortening URL",
    })
}
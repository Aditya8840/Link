package controllers

import (
	"net/http"

	"github.com/Aditya8840/Link/types"
	"github.com/gofiber/fiber/v2"
)

func ShortURL(c *fiber.Ctx) error {
	var url types.URL
	err := c.BodyParser(&url)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
	}
	return c.JSON(fiber.Map{
        "message": "Shortening URL",
    })
}
package controllers

import (
	"github.com/Aditya8840/Link/databases"
	"github.com/gofiber/fiber/v2"
)

func RedirectionURL(c *fiber.Ctx) error {
	code := c.Params("code")

	url, err := databases.Mgr.GetOriginalURL(code)
	if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Short URL not found",
        })
    }

	return c.Redirect(url, fiber.StatusPermanentRedirect)

}
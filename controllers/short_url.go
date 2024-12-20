package controllers

import (
	"net/http"
	"time"

	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/databases"
	"github.com/Aditya8840/Link/types"
	"github.com/Aditya8840/Link/utils"
	"github.com/gofiber/fiber/v2"
)

func ShortURL(c *fiber.Ctx) error {
	var url types.ShortURL
	err := c.BodyParser(&url)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
	}

	count, err:= databases.Mgr.GetAndIncCounter()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
	}
	urlCode := utils.Base62Encode(count)
	var urlObject types.URL

	urlObject.CreatedAt = time.Now().Unix()
	urlObject.URLCode = urlCode
	urlObject.LongURL = url.LongURL

	err = databases.Mgr.Insert(&urlObject)
	if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
	return c.Status(http.StatusOK).JSON(
		fiber.Map{
            "short_url": constant.BASE_URL+urlCode,
        },
	)
}
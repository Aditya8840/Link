package routes

import (
	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/controllers"
	"github.com/gofiber/fiber/v2"
)

func ShortRoutes(app *fiber.App) {
	app.Post(constant.ShortURLPath, controllers.ShortURL)
}
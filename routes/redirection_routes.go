package routes

import (
	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/controllers"
	"github.com/gofiber/fiber/v2"
)

func RedirectRoutes(app *fiber.App){
	app.Get(constant.RedirectionURLPath, controllers.RedirectionURL)
}
package main

import (
	"fmt"

	"github.com/Aditya8840/Link/databases"
	"github.com/Aditya8840/Link/routes"
	"github.com/gofiber/fiber/v2"
)


func main() {
	fmt.Println("Hello, World!")

	app := fiber.New()

	databases.Connect()
	routes.RedirectRoutes(app)
	routes.ShortRoutes(app)

	app.Listen(":8080")
}

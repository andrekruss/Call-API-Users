package main

import (
	"call-ms-users/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", controllers.GetRoot)
	app.Post("/create-user", controllers.PostUser)

	app.Listen(":3000")
}

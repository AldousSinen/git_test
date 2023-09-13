package main

import (
	"sinen_project/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/test", controller.GetUser)
	app.Get("/test/:firstname", controller.ShowUser)

	app.Listen(":3030")
}

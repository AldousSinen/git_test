package controller

import (
	"sinen_project/model"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	requestBody := &model.TestFields{}
	c.BodyParser(requestBody)
	return c.JSON(requestBody)
}

func ShowUser(c *fiber.Ctx) error {
	firstname := c.Params("firstname")
	return c.SendString(firstname)
}

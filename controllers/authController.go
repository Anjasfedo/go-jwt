package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	c.BodyParser()
	return c.SendString("Hello, World 👋!")
}

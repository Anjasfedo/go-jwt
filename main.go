package main

import (
	"github.com/Anjasfedo/go-react-jwt/database"
	"github.com/Anjasfedo/go-react-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}

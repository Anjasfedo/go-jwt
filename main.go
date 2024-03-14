package main

import (
	"github.com/Anjasfedo/go-react-jwt/database"
	"github.com/Anjasfedo/go-react-jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		// AllowOriginsFunc: func(origin string) bool { return true },
		AllowOriginsFunc: func(origin string) bool {
            return strings.Contains(origin, ":://localhost")
        },
	  }))

	routes.Setup(app)

	app.Listen(":3000")
}

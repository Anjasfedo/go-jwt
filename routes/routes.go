package routes

import (
	"github.com/Anjasfedo/go-react-jwt/controllers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
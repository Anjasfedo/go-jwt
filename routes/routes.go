package routes

import (
	"github.com/Anjasfedo/go-react-jwt/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
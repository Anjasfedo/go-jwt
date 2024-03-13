package controllers

import (
	"github.com/Anjasfedo/go-react-jwt/database"
	"github.com/Anjasfedo/go-react-jwt/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	bytePassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	password := string(bytePassword)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

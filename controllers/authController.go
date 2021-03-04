package controllers

import (
	"github.com/al-abbas-nz/golang-react/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	return c.JSON(user)
}

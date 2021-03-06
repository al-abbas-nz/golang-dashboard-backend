package controllers

import (
	"github.com/al-abbas-nz/golang-react/database"
	"github.com/al-abbas-nz/golang-react/middleware"
	"github.com/al-abbas-nz/golang-react/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "Permissions"); err != nil {
		return err
	}
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}

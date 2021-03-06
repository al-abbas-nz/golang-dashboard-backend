package controllers

import (
	"strconv"

	"github.com/al-abbas-nz/golang-react/database"
	"github.com/al-abbas-nz/golang-react/models"
	"github.com/gofiber/fiber/v2"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}

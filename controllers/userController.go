package controllers

import (
	"strconv"

	"github.com/al-abbas-nz/golang-react/database"
	"github.com/al-abbas-nz/golang-react/middleware"
	"github.com/al-abbas-nz/golang-react/models"
	"github.com/gofiber/fiber/v2"
)

// CRUD for users

// get all users (objects in an array)
func AllUsers(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	// add pagination
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

// create user
func CreateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

// get a user
func GetUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Find(&user)

	return c.JSON(user)
}

// update user
func UpdateUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}

package main

import (
	"github.com/al-abbas-nz/golang-react/database"
	"github.com/al-abbas-nz/golang-react/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}

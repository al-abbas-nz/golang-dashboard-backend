package routes

import (
	"github.com/al-abbas-nz/golang-react/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}

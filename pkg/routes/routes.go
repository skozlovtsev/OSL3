package routes

import (
	"osl3/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	app.Get("/login", controller.LogInPage)
	app.Get("/validate", controller.LogIn)
	app.Get("/check", controller.CheckUser)
}

package main

import (
	"osl3/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.AddRoutes(app)

	panic(app.Listen(":7999"))
}

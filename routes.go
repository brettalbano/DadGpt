package main

import (
	"github.com/brettalbano/DadGpt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.UserIndex)
}

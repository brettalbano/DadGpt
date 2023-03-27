package main

import (
	"github.com/brettalbano/DadGpt/controllers"
	"github.com/brettalbano/DadGpt/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.UserIndex)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
	app.Get("/validate", middleware.RequireAuth, middleware.ExtendSession, controllers.ValidateUser)
	app.Post("/logout", middleware.RequireAuth, controllers.LogoutUser)
	app.Get("/users", controllers.GetUsers)
}

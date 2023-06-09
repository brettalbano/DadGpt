package main

import (
	"DadGpt/controllers"
	"DadGpt/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.UserIndex)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
	app.Get("/validate", middleware.RequireAuth, controllers.ValidateUser)
	app.Post("/logout", middleware.RequireAuth, controllers.LogoutUser)
	app.Get("/users", controllers.GetUsers)
	app.Post("/conversation", controllers.HaveConversation)
}

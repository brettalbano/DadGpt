package main

import (
	"os"

	"github.com/brettalbano/DadGpt/initializers"
	"github.com/brettalbano/DadGpt/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	// Load templates
	engine := html.New("./views", ".tmpl")

	// Setup app.

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")
	app.Use(middleware.RequireAuth)

	// Routes.
	Routes(app)

	// Start App.
	app.Listen(":" + os.Getenv("PORT"))

}

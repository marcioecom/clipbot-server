package api

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", healthCheck)

	api.Post("/upload", handleFileUpload)
}

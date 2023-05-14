package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcioecom/clipbot-server/api/user"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", healthCheck)

	userController := user.NewUserController()
	api.Get("/users", userController.GetAll)
	api.Post("/users", userController.Create)
}

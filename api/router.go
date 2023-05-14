package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcioecom/clipbot-server/api/auth"
	"github.com/marcioecom/clipbot-server/api/upload"
	"github.com/marcioecom/clipbot-server/api/user"
	"github.com/marcioecom/clipbot-server/infra/storage"
	"go.uber.org/zap"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Health
	api.Get("/health", healthCheck)

	// User
	u := api.Group("/users")
	userController := user.NewController()
	u.Get("/", userController.GetAll)
	u.Post("/", userController.Create)

	// Auth
	a := api.Group("/auth")
	authController := auth.NewController()
	a.Post("/", authController.Login)

	// Upload
	p := api.Group("/upload")

	s, err := storage.New()
	if err != nil {
		zap.L().Fatal("failed to start storage", zap.Error(err))
	}

	uploadController := upload.NewController(s)
	p.Post("/", protected(), uploadController.Save)
}

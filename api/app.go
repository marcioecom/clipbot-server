package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/marcioecom/clipbot-server/helper"
	"go.uber.org/zap"
)

// Start starts the API server
func Start() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${ip}:${port} ${time} ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))

	app.Use(cors.New())
	app.Static("/videos", "./videos")

	setupRoutes(app)

	port := fmt.Sprintf(":%s", helper.GetEnv("PORT"))
	if err := app.Listen(port); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}

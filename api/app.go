package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Setup starts the API server
func Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 2000,
	})

	app.Use(logger.New(logger.Config{
		Format:     "${ip}:${port} ${time} ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	app.Static("/videos", "./videos")

	setupRoutes(app)

	return app
}

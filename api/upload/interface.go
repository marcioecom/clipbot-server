package upload

import "github.com/gofiber/fiber/v2"

type IUploadController interface {
	Save(c *fiber.Ctx) error
}

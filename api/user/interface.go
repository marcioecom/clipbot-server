package user

import "github.com/gofiber/fiber/v2"

type IUserController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

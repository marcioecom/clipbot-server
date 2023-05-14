package auth

import "github.com/gofiber/fiber/v2"

type IAuthController interface {
	Login(ctx *fiber.Ctx) error
}

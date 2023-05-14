package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/marcioecom/clipbot-server/helper"
	"github.com/marcioecom/clipbot-server/infra/database"
	"github.com/marcioecom/clipbot-server/infra/database/models"
	"github.com/marcioecom/clipbot-server/infra/database/repositories"
)

type controller struct {
	repository repositories.IUserRepository
}

func NewUserController() IUserController {
	return &controller{
		repository: repositories.NewUserRepository(database.DB),
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	u := new(models.Users)

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := validator.New().Struct(u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"error":   helper.HandleValidatorErr(err),
		})
	}

	user, err := c.repository.GetByEmail(u.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get user",
		})
	}

	if user != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "User already exists",
		})
	}

	if err := c.repository.Create(u); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    u,
	})
}

func (c *controller) GetAll(ctx *fiber.Ctx) error {
	users, err := c.repository.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get users",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

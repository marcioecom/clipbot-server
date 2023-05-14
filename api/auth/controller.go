package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/marcioecom/clipbot-server/helper"
	"github.com/marcioecom/clipbot-server/infra/database"
	"github.com/marcioecom/clipbot-server/infra/database/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthController struct {
	userRepo repositories.IUserRepository
}

func NewController() IAuthController {
	return &AuthController{
		userRepo: repositories.NewUserRepository(database.DB),
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := new(LoginInput)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"error":   err.Error(),
		})
	}

	if err := validator.New().Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"error":   helper.HandleValidatorErr(err),
		})
	}

	u, err := c.userRepo.GetByEmail(input.Email)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on email",
			"error":   err.Error(),
		})
	}

	if u == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	if !validPasswordHash(input.Password, u.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(helper.GetEnv("SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data":    t,
	})
}

// validPasswordHash compare password with hash
func validPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

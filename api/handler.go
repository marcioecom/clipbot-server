package api

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/marcioecom/clipbot-server/constants"
	"github.com/marcioecom/clipbot-server/helper"
	"github.com/marcioecom/clipbot-server/infra/queue"
	"go.uber.org/zap"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Server is healthy",
	})
}

func handleFileUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("video upload error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error uploading video",
		})
	}

	id := uuid.New()

	filename := strings.Replace(id.String(), "-", "", -1)
	fileext := strings.Split(file.Filename, ".")[1]

	video := fmt.Sprintf("%s.%s", filename, fileext)

	// TODO: upload to a cloud storage
	if err := c.SaveFile(file, fmt.Sprintf("./videos/%s", video)); err != nil {
		zap.L().Error("video save error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error saving video",
		})
	}

	if err := queue.Producer.Produce(constants.ClipTopic, []byte(video)); err != nil {
		zap.L().Error("failed to send message", zap.Error(err))
	}

	videourl := fmt.Sprintf("%s/videos/%s", helper.GetEnv("host"), video)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":   true,
		"message":   "Video uploaded successfully",
		"videoUrl":  videourl,
		"videoName": video,
		"header":    file.Header,
		"size":      file.Size,
	})
}

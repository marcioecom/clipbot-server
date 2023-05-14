package upload

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/marcioecom/clipbot-server/constants"
	"github.com/marcioecom/clipbot-server/helper"
	"github.com/marcioecom/clipbot-server/infra/queue"
	"github.com/marcioecom/clipbot-server/infra/storage"
	"go.uber.org/zap"
)

type UploadController struct {
	storage storage.IStorage
}

func NewController(s storage.IStorage) IUploadController {
	return &UploadController{
		storage: s,
	}
}

func (u *UploadController) Save(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("video upload error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("error uploading video: %v", err),
		})
	}

	id := uuid.New()

	filename := strings.Replace(id.String(), "-", "", -1)
	fileext := filepath.Ext(file.Filename)

	video := filename + fileext
	uploadDir := fmt.Sprintf("../videos/%s", video)

	body, err := file.Open()
	if err != nil {
		zap.L().Error("video open error", zap.Error(err), zap.String("video", video), zap.String("uploadDir", uploadDir))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("error opening video: %v", err),
		})
	}

	if err := u.storage.Upload(video, body); err != nil {
		zap.L().Error("video save error", zap.Error(err), zap.String("video", video), zap.String("uploadDir", uploadDir))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("error saving video: %v", err),
		})
	}

	if err := queue.Producer.Produce(
		helper.GetEnv("QUEUE_TOPIC").FallBack(constants.ClipTopic),
		[]byte(video),
	); err != nil {
		zap.L().Error("failed to send message", zap.Error(err))
	}

	videourl := fmt.Sprintf("%s/videos/%s", helper.GetEnv("host"), video)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"message":   "video uploaded successfully",
		"videoUrl":  videourl,
		"videoName": video,
		"size":      file.Size,
	})
}

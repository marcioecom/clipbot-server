package helper

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes the logger based on the environment
func InitLogger() {
	var logconfig zap.Config

	if strings.EqualFold(os.Getenv("ENV"), "production") {
		logconfig = zap.NewProductionConfig()
	} else {
		logconfig = zap.NewDevelopmentConfig()
	}

	logconfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := logconfig.Build()
	if err != nil {
		logger.Fatal("failed to initialize logger", zap.Error(err))
	}

	zap.ReplaceGlobals(logger)
}

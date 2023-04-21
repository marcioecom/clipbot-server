package helper

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes the logger based on the environment
func InitLogger() {
	logconfig := zap.NewDevelopmentConfig()
	if strings.EqualFold(GetEnv("ENV").String(), "production") {
		logconfig = zap.NewProductionConfig()
	}

	logconfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := logconfig.Build()
	if err != nil {
		zap.L().Fatal("failed to initialize logger", zap.Error(err))
	}

	zap.ReplaceGlobals(logger)
}

package helper

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// MemEnvs is a map of envs to avoid make system calls
// preset values are used if the env is not set
var MemEnvs = map[string]string{
	"PORT": "3001",
}

// LoadEnvs loads envs from the system
func LoadEnvs() {
	for key := range MemEnvs {
		env := os.Getenv(key)
		if env == "" {
			zap.L().Info("using default", zap.String("key", key), zap.String("value", env))
			continue
		}
		MemEnvs[key] = env
	}
}

// GetEnv returns the value of a given env
func GetEnv(key string) string {
	value, ok := MemEnvs[strings.ToUpper(key)]
	if !ok {
		zap.L().Fatal("env not found", zap.String("key", key))
	}
	return value
}

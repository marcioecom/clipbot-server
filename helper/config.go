package helper

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// MemEnvs is a map of envs to avoid make system calls
// preset values are used if the env is not set
var MemEnvs = map[string]string{
	"PORT":           "3001",
	"HOST":           "http://localhost:3001",
	"KAFKA_URL":      "localhost:9092",
	"KAFKA_USERNAME": "admin",
	"KAFKA_PASSWORD": "admin",
}

// LoadEnvs loads envs from the system
func LoadEnvs() {
	for key := range MemEnvs {
		env, ok := os.LookupEnv(key)
		if !ok {
			zap.L().Info("using default env", zap.String("key", key), zap.String("value", MemEnvs[key]))
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

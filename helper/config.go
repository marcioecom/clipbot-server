package helper

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// MemEnvs is a map of envs to avoid make system calls
// preset values are used if the env is not set
var MemEnvs = map[string]string{
	"ENV":    "DEVELOP",
	"PORT":   "3001",
	"SECRET": "secret",
	"HOST":   "http://localhost:3001",

	"KAFKA_URL":      "localhost:9092",
	"KAFKA_USERNAME": "admin",
	"KAFKA_PASSWORD": "admin",
	"QUEUE_TOPIC":    "clips.send-clip",

	"DB_HOST":     "",
	"DB_PORT":     "",
	"DB_NAME":     "",
	"DB_USER":     "",
	"DB_PASSWORD": "",

	"AWS_KEY":      "",
	"AWS_SECRET":   "",
	"AWS_REGION":   "",
	"AWS_BUCKET":   "",
	"AWS_ENDPOINT": "",
}

// LoadEnvs loads envs from the system
func LoadEnvs() {
	if err := godotenv.Load(); err != nil {
		zap.L().Warn("no .env file found")
	}

	for key := range MemEnvs {
		env, ok := os.LookupEnv(key)
		if !ok {
			zap.L().Info("using default env", zap.String("key", key), zap.String("value", MemEnvs[key]))
			continue
		}
		MemEnvs[key] = env
	}
}

type Env string

// GetEnv returns the value of a given env
func GetEnv(key string) Env {
	value, ok := MemEnvs[strings.ToUpper(key)]
	if !ok {
		zap.L().Warn("env not found in MemEnvs", zap.String("key", key))
	}
	return Env(value)
}

func (e Env) FallBack(fallback string) string {
	if e != "" {
		return e.String()
	}
	return fallback
}

func (e Env) String() string {
	return string(e)
}

func (e Env) Bool() bool {
	return strings.EqualFold(string(e), "true")
}

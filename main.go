package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/marcioecom/clipbot-server/api"
	"github.com/marcioecom/clipbot-server/helper"
	"github.com/marcioecom/clipbot-server/infra/queue"
	"go.uber.org/zap"
)

func main() {
	helper.InitLogger()
	helper.LoadEnvs()

	queue.Start(&queue.Config{
		GroupID:      "my-group",
		QueueURL:     helper.GetEnv("kafka_url"),
		SaslUsername: helper.GetEnv("kafka_username"),
		SaslPassword: helper.GetEnv("kafka_password"),
	})

	app := api.Setup()
	port := fmt.Sprintf(":%s", helper.GetEnv("PORT"))
	if err := app.Listen(port); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	zap.L().Info("stopping service")
	go func() {
		time.Sleep(3 * time.Second)
		zap.L().Panic("stop timeout")
	}()

	if err := queue.Consumer.Stop(); err != nil {
		zap.L().Error("failed to stop consumer", zap.Error(err))
	}
}

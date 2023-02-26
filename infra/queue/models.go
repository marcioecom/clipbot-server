package queue

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type consumer struct {
	client *kafka.Consumer
}

type producer struct {
	client *kafka.Producer
}

// Config is the configuration for the queue
type Config struct {
	QueueURL     string `validate:"required"`
	GroupID      string `validate:"required"`
	SaslUsername string `validate:"required"`
	SaslPassword string `validate:"required"`
}

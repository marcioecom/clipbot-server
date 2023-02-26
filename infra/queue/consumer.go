package queue

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

// NewConsumer creates a new consumer instance
func NewConsumer(config *Config) (IConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.QueueURL,
		"group.id":          config.GroupID,
		"sasl.username":     config.SaslUsername,
		"sasl.password":     config.SaslPassword,
		"auto.offset.reset": "earliest",
		"sasl.mechanism":    "SCRAM-SHA-256",
		"security.protocol": "SASL_SSL",
	})
	if err != nil {
		zap.L().Debug("could not init consumer", zap.Error(err))
		return nil, err
	}

	zap.L().Debug("consumer started")
	return &consumer{
		client: c,
	}, nil
}

func (c *consumer) Consume(topic string, handler func(*kafka.Consumer, kafka.Event) error) error {
	return c.client.Subscribe(topic, handler)
}

func (c *consumer) Stop() error {
	return c.client.Close()
}

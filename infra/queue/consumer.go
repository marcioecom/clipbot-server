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
		zap.L().Error("could not init consumer", zap.Error(err))
		return nil, err
	}

	zap.L().Debug("consumer started")
	return &consumer{
		client: c,
	}, nil
}

func (c *consumer) Consume(topic string, handler func(*kafka.Consumer, kafka.Event) error) error {
	if err := c.client.Subscribe(topic, nil); err != nil {
		return err
	}

	for {
		ev := c.client.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				zap.L().Error("kafka error", zap.Error(e.TopicPartition.Error))
				continue
			}

			if err := handler(c.client, e); err != nil {
				zap.L().Error("could not handle message", zap.Error(err))
			}
		case kafka.Error:
			zap.L().Error("kafka error", zap.Error(e))
		}
	}
}

func (c *consumer) Stop() error {
	return c.client.Close()
}

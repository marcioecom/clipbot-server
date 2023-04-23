package queue

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

// NewProducer creates a new producer instance
func NewProducer(config *Config) (IProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.QueueURL,
		"sasl.username":     config.SaslUsername,
		"sasl.password":     config.SaslPassword,
		"sasl.mechanism":    "SCRAM-SHA-256",
		"security.protocol": "SASL_SSL",
	})
	if err != nil {
		return nil, err
	}

	zap.L().Info("producer started")
	return &producer{
		client: p,
	}, nil
}

func (p *producer) Produce(topic string, data []byte) error {
	return p.client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
	}, nil)
}

func (p *producer) Stop() {
	p.client.Close()
}

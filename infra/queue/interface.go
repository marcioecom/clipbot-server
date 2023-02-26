package queue

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

// IConsumer define the consumer methods
type IConsumer interface {
	Consume(topic string, handler func(*kafka.Consumer, kafka.Event) error) error
	Stop() error
}

// IProducer define the producer methods
type IProducer interface {
	Produce(topic string, data []byte) error
	Stop()
}

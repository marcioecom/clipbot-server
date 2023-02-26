package queue

import (
	"github.com/go-playground/validator/v10"
)

var (
	// Consumer is the consumer instance
	Consumer IConsumer
	// Producer is the producer instance
	Producer IProducer
)

// Start starts producer and consumer
func Start(config *Config) error {
	err := validator.New().Struct(config)
	if err != nil {
		return err
	}

	Producer, err = NewProducer(config)
	if err != nil {
		return err
	}

	Consumer, err = NewConsumer(config)
	if err != nil {
		return err
	}

	return nil
}

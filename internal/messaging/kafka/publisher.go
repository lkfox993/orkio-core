package kafka

import (
	Kafka "github.com/segmentio/kafka-go"
)

type Publisher struct {
	publisher *Kafka.Writer
}

func NewPublisher(publisher *Kafka.Writer) *Publisher {
	return &Publisher{
		publisher: publisher,
	}
}

func (p *Publisher) Publish() error {
	return nil
}

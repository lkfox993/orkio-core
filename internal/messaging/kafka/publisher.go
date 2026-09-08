package kafka

import (
	"context"

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

func (p *Publisher) Publish(ctx context.Context, data []byte, topic string) error {

	msg := Kafka.Message{Topic: topic, Value: data}
	return p.publisher.WriteMessages(ctx, msg)
}

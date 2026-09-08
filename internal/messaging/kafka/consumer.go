package kafka

import (
	"context"

	"github.com/lkfox993/orkio-core/internal/messaging"
	Kafka "github.com/segmentio/kafka-go"
)

type Consumer struct {
	consumer *Kafka.Reader
}

func NewConsumer(consumer *Kafka.Reader) *Consumer {
	return &Consumer{
		consumer: consumer,
	}
}

func (c *Consumer) Consume(ctx context.Context, handler messaging.Handler) error {

	for {
		msg, err := c.consumer.FetchMessage(ctx)

		if err != nil {
			return err
		}

		if err := handler(ctx, msg.Value); err != nil {
			continue
		}

		if err := c.consumer.CommitMessages(ctx, msg); err != nil {
			return err
		}
	}

}

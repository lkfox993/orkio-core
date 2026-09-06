package pulsar

import (
	"context"

	Pulsar "github.com/apache/pulsar-client-go/pulsar"
	"github.com/lkfox993/orkio-core/internal/messaging"
)

type Consumer struct {
	consumer Pulsar.Consumer
}

func NewConsumer(consumer Pulsar.Consumer) *Consumer {
	return &Consumer{
		consumer: consumer,
	}
}

func (c *Consumer) Consume(ctx context.Context, handler messaging.Handler) error {

	for {
		msg, err := c.consumer.Receive(ctx)

		if err != nil {
			return err
		}

		if err := handler(ctx, msg.Payload()); err != nil {
			c.consumer.Nack(msg)
			continue
		}

		c.consumer.Ack(msg)
	}

}

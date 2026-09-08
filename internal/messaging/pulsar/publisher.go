package pulsar

import (
	"context"

	Pulsar "github.com/apache/pulsar-client-go/pulsar"
)

type Publisher struct {
	publisher Pulsar.Producer
}

func NewPublisher(publisher Pulsar.Producer) *Publisher {
	return &Publisher{
		publisher: publisher,
	}
}

func (p *Publisher) Publish(ctx context.Context, data []byte, topic string) error {

	msg := Pulsar.ProducerMessage{
		Payload: data,
	}

	_, err := p.publisher.Send(ctx, &msg)

	return err
}

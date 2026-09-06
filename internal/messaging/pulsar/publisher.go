package pulsar

import (
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

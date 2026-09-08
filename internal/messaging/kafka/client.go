package kafka

import (
	"time"

	Kafka "github.com/segmentio/kafka-go"
)

type Client struct {
	dialer  *Kafka.Dialer
	brokers []string
}

func NewClient(brokers []string) *Client {

	dialer := &Kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	return &Client{
		dialer:  dialer,
		brokers: brokers,
	}
}

func (c *Client) NewConsumer(
	topic string,
	groupID string,
) *Consumer {

	r := Kafka.NewReader(Kafka.ReaderConfig{
		Brokers: c.brokers,
		Topic:   topic,
		GroupID: groupID,
		Dialer:  c.dialer,
	})

	return &Consumer{
		consumer: r,
	}
}

func (c *Client) NewPublisher(topic string) *Publisher {

	w := Kafka.NewWriter(Kafka.WriterConfig{
		Brokers:  c.brokers,
		Topic:    topic,
		Balancer: &Kafka.LeastBytes{},
		Dialer:   c.dialer,
	})

	return &Publisher{
		publisher: w,
	}
}

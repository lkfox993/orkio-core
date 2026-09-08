package pulsar

import (
	"time"

	Pulsar "github.com/apache/pulsar-client-go/pulsar"
)

type Client struct {
	conn Pulsar.Client
}

func NewClient(url string) (Pulsar.Client, error) {
	return Pulsar.NewClient(Pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
}

func (c *Client) NewPublisher(topic string) (*Publisher, error) {

	r, err := c.conn.CreateProducer(Pulsar.ProducerOptions{
		Topic: topic,
	})

	return &Publisher{
		publisher: r,
	}, err
}

func (c *Client) NewConsumer(topic string) (*Consumer, error) {

	w, err := c.conn.Subscribe(Pulsar.ConsumerOptions{
		Topic: topic,
	})

	return &Consumer{
		consumer: w,
	}, err
}

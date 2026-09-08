package pulsar

import (
	"time"

	Pulsar "github.com/apache/pulsar-client-go/pulsar"
)

func NewClient(url string) (Pulsar.Client, error) {
	return Pulsar.NewClient(Pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
}

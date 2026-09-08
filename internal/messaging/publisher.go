package messaging

import "context"

type Publisher interface {
	Publish(ctx context.Context, data []byte, topic string) error
}

package messaging

import "context"

type Handler func(ctx context.Context, data []byte) error

type Consumer interface {
	Consume(ctx context.Context, handler Handler) error
}

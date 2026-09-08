package coordinator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkfox993/orkio-core/internal/events"
	"github.com/lkfox993/orkio-core/internal/messaging"
)

type Consumer struct {
	engine   *Engine
	consumer messaging.Consumer
}

func NewConsumer(engine *Engine, consumer messaging.Consumer) *Consumer {
	return &Consumer{
		engine:   engine,
		consumer: consumer,
	}
}

func (c *Consumer) HandleMessage(ctx context.Context, data []byte) error {

	var event events.Event

	if err := json.Unmarshal(data, &event); err != nil {
		return fmt.Errorf("unmarshal event: %w", err)
	}

	switch event.Type {
	case events.TypeJobCompleted:

		var payload events.JobCompleted

		if err := json.Unmarshal(event.Data, &payload); err != nil {
			return fmt.Errorf("unmarshal job.completed: %w", err)
		}

		return c.engine.HandleJobCompleted(ctx, payload)

	case events.TypeJobFailed:

		var payload events.JobFailed

		if err := json.Unmarshal(event.Data, &payload); err != nil {
			return fmt.Errorf("unmarshal job.failed: %w", err)
		}

		return c.engine.HandleJobFailed(ctx, payload)

	case events.TypeJobTimedOut:

		var payload events.JobTimedOut

		if err := json.Unmarshal(event.Data, &payload); err != nil {
			return fmt.Errorf("unmarshal job.timed_out: %w", err)
		}

		return c.engine.HandleJobTimedOut(ctx, payload)

	// case events.TypeWorkflowStarted:
	// 	return c.engine.HandleWorkflowStarted(ctx, event)

	default:
		return fmt.Errorf("unknown event type: %q", event.Type)
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	return c.consumer.Consume(ctx, c.HandleMessage)
}

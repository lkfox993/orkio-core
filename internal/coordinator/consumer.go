package coordinator

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lkfox993/orkio-core/internal/events"
)

type Consumer struct {
	engine *Engine
}

func NewConsumer(engine *Engine) *Consumer {
	return &Consumer{
		engine: engine,
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

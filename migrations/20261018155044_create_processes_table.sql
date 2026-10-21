-- +goose Up
CREATE TABLE processes (
    id UUID PRIMARY KEY,
    workflow_id UUID NOT NULL,
    status VARCHAR(50) NOT NULL,
    initial_variables JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE processes;
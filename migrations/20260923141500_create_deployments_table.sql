-- +goose Up
CREATE TABLE deployments (
    id          UUID PRIMARY KEY,
    tenant_id   UUID NOT NULL,
    name        VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    metadata    JSONB
);

CREATE INDEX idx_deployments_tenant_id ON deployments (tenant_id);

-- +goose Down
DROP TABLE deployments;

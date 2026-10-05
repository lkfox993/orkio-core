-- +goose Up
CREATE TABLE workflows (

    id              UUID PRIMARY KEY,
    tenant_id       UUID NOT NULL,
    key             VARCHAR(255) NOT NULL,
    name            VARCHAR(255),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP

);

-- +goose Down
DROP TABLE workflows;

-- +goose Up
CREATE TABLE workflow_versions (
    
    id              UUID PRIMARY KEY,
    tenant_id       UUID NOT NULL,
    workflow_id     UUID NOT NULL REFERENCES workflows(id),
    deployment_id   UUID NOT NULL REFERENCES deployments(id),
    version         INTEGER NOT NULL,
    resource_name   VARCHAR(255),
    definition      TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (workflow_id, version)
);

-- +goose Down
DROP TABLE workflow_versions;
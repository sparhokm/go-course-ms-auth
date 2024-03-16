-- +goose Up
INSERT INTO endpoint_accesses (endpoint, role, updated_at, created_at) VALUES
    ('ChatV1.Create', 'admin', now(), now());

-- +goose Down
truncate table endpoint_accesses;
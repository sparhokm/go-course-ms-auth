-- +goose Up
create table endpoint_accesses
(
    id         serial primary key,
    endpoint   varchar(255),
    role       varchar(64),
    created_at timestamp not null default now(),
    updated_at timestamp,
    UNIQUE (endpoint, role)
);

-- +goose Down
drop table endpoint_accesses;

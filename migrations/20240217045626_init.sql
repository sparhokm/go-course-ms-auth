-- +goose Up
create table users
(
    id            serial primary key,
    name          varchar(255) not null,
    email         varchar(255) not null,
    password_hash varchar(255) not null,
    role          varchar(64)  not null,
    created_at    timestamp    not null default now(),
    updated_at    timestamp
);

-- +goose Down
drop table users;
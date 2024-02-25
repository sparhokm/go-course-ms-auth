-- +goose Up
create table deleted_users
(
    id         serial primary key,
    deleted_at timestamp not null default now()
);

-- +goose Down
drop table deleted_users;
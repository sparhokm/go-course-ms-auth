-- +goose Up
INSERT INTO endpoint_accesses (endpoint, role, updated_at, created_at)
VALUES ('/chat_v1.ChatV1/Create', 'admin', now(), now()),
       ('/chat_v1.ChatV1/Delete', 'admin', now(), now()),
       ('/chat_v1.ChatV1/SendMessage', 'admin', now(), now()),
       ('/chat_v1.ChatV1/SendMessage', 'user', now(), now()),
       ('/chat_v1.ChatV1/ConnectChat', 'user', now(), now());

-- +goose Down
truncate table endpoint_accesses;
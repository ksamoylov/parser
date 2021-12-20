CREATE TABLE users
(
    id         BIGSERIAL primary key,
    username TEXT not null,
    name  TEXT,
    email VARCHAR,
    updated_at TIMESTAMP default now(),
    created_at TIMESTAMP default now()
);
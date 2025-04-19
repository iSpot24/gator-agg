-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    name TEXT unique not null, 
    created_at TIMESTAMP default CURRENT_TIMESTAMP not null, 
    updated_at TIMESTAMP default CURRENT_TIMESTAMP not null
    );

-- +goose Down
DROP TABLE users;
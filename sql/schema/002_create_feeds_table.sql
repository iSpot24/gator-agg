-- +goose Up
CREATE TABLE IF NOT EXISTS feeds(
    id INTEGER PRIMARY KEY generated always as identity,
    name TEXT not null,
    url TEXT UNIQUE not null,
    user_id UUID not null REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP default CURRENT_TIMESTAMP not null, 
    updated_at TIMESTAMP default CURRENT_TIMESTAMP not null
    );

-- +goose Down
DROP TABLE feeds;
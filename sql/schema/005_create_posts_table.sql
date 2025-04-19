-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY generated always as identity,
    title TEXT,
    url TEXT,
    description TEXT,
    feed_id INTEGER NOT NULL REFERENCES feeds(id),
    published_at TIMESTAMP,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP default CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE posts;
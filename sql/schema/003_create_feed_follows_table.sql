-- +goose Up
CREATE TABLE IF NOT EXISTS feed_follows(
    id INTEGER PRIMARY KEY generated always as identity,
    user_id UUID not null REFERENCES users(id) ON DELETE CASCADE,
    feed_id INTEGER not null REFERENCES feeds(id) ON DELETE CASCADE,
    created_at TIMESTAMP default CURRENT_TIMESTAMP not null, 
    updated_at TIMESTAMP default CURRENT_TIMESTAMP not null,
    UNIQUE(user_id, feed_id)
    );

-- +goose Down
DROP TABLE feed_follows;
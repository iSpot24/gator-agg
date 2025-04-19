-- name: GetFeeds :many
SELECT * FROM feeds f
INNER JOIN users u on u.id = f.user_id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds f where f.url = $1 LIMIT 1;

-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
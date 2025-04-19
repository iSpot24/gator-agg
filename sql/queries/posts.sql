-- name: CreatePost :one
INSERT INTO posts (title, url, description, feed_id, published_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetPostsByUser :many
SELECT p.* FROM posts p
INNER JOIN feed_follows ff on ff.feed_id = p.feed_id
WHERE ff.user_id = $1
ORDER BY p.created_at DESC
LIMIT $2;
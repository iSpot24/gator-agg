-- name: CreateFeedFollow :one
WITH inserted as (
    INSERT INTO feed_follows (user_id, feed_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4
    )
    RETURNING *
) SELECT i.*, u.name as user_name, f.name as feed_name FROM inserted i
INNER JOIN users u on u.id = i.user_id
INNER JOIN feeds f on f.id = i.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT *, u.name as user_name FROM feeds f
INNER JOIN feed_follows ff on f.id = ff.feed_id
INNER JOIN users u on u.id = ff.user_id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollowByUserAndFeed :exec
DELETE FROM feed_follows 
WHERE feed_follows.id IN (
    SELECT ff.id FROM feed_follows ff
    INNER JOIN users u on u.id = ff.user_id
    INNER JOIN feeds f on f.id = ff.feed_id
    WHERE u.id = $1 and f.url = $2
    );
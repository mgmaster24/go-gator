
-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
      $1,
      $2,
      $3,
      $4,
      $5
    )
  RETURNING *)
SELECT 
  inserted_feed_follow.*,
  feeds.name as feed_name,
  users.name as user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users on inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowById :one
SELECT * FROM feed_follows WHERE id = $1;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, users.name as user_name, feeds.name as feed_name 
FROM feed_follows 
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;


-- name: GetFeedFollows :many
SELECT id, feed_id, user_id FROM feed_follows;

-- name: DeleteFeedFollow :execresult
DELETE 
FROM feed_follows
USING feeds 
WHERE feed_follows.user_id = $1 AND feeds.url = $2;
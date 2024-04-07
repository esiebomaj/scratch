-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetEarliestFetchedFeeds :many
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: UpdateLastFetchedAt :one
UPDATE feeds
SET last_fetched_at = $1, updated_at = $1
WHERE id = $2
RETURNING *;


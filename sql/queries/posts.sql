-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, link, description, feed_id, published_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetAllPosts :many
SELECT * FROM posts;

-- name: GePostForFeed :many
SELECT * FROM posts
WHERE feed_id = $1
ORDER BY created_at DESC
LIMIT 10;


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

-- name: GetRecentPostForUser :many
SELECT p.* FROM posts p
JOIN feed_follows ff ON p.feed_id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY p.created_at DESC
LIMIT $2;


-- name: CheckPostExist :one
SELECT EXISTS(SELECT 1 FROM posts WHERE link = $1);

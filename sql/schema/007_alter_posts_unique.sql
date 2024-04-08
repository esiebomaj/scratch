-- +goose Up
ALTER TABLE posts 
ADD CONSTRAINT posts_link_unique UNIQUE(link);

-- +goose Down
ALTER TABLE posts
DROP CONSTRAINT posts_link_unique;
-- +goose Up
ALTER TABLE feed_follows 
ADD CONSTRAINT feed_id_user_id_unique UNIQUE(user_id, feed_id);

-- +goose Down
ALTER TABLE feed_follows
DROP CONSTRAINT feed_id_user_id_unique;
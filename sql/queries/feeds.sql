-- name: CreateFeed :one
INSERT INTO feeds (user_id, name, url)
VALUES (
	(SELECT id FROM users WHERE id=$1),
	$2,
	$3
)
RETURNING *; 

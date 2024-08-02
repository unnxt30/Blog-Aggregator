-- name: GetPostsByUsers :many
Select * from posts ORDER BY published_at LIMIT $1;

-- name: CreatePost :one
INSERt INTO posts(id, created_at, updated_at, title,  url, description, published_at, feed_id)
values($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserPost :many
select * from posts where feed_id = $1 LIMIT $2;


-- name: CreateFeed :one
INSERT into feed (feed_id, created_at, updated_at, name, url, user_id)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :many
Select * from feed;

-- name: GetNextFeedsToFetch :many
Select * from feed 
ORDER BY last_fetched_at NULLS first LIMIT $1;

-- name: MarkFeedFetched :exec
Update feed 
set last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
where feed_id = $1;
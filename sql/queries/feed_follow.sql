-- name: CreateFeedfollow :one 
insert into feed_follow(feed_follow_id, created_at, updated_at, user_id, feed_id)
values($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedfollow :exec
Delete from feed_follow 
where feed_follow_id = $1;
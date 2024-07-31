// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedfollow = `-- name: CreateFeedfollow :one
insert into feed_follow(feed_follow_id, created_at, updated_at, user_id, feed_id)
values($1, $2, $3, $4, $5)
RETURNING feed_follow_id, created_at, updated_at, user_id, feed_id
`

type CreateFeedfollowParams struct {
	FeedFollowID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uuid.UUID
	FeedID       uuid.UUID
}

func (q *Queries) CreateFeedfollow(ctx context.Context, arg CreateFeedfollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedfollow,
		arg.FeedFollowID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.FeedFollowID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFeedfollow = `-- name: DeleteFeedfollow :exec
Delete from feed_follow 
where feed_follow_id = $1
`

func (q *Queries) DeleteFeedfollow(ctx context.Context, feedFollowID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedfollow, feedFollowID)
	return err
}
-- +goose Up
Create table feed_follow(
    feed_follow_id UUID,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    user_id UUID not null,
    feed_id UUID not null,
    PRIMARY KEY(feed_follow_id),
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(feed_id) REFERENCES feed(feed_id)
);

-- +goose Down

Drop table feed_follow;

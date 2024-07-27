-- +goose Up
CREATE TABLE feed(
    feed_id UUID not null,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null, 
    name varchar(30), 
    url varchar(60) UNIQUE,
    user_id UUID not null ,
    PRIMARY KEY(feed_id), 
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE);

-- +goose Down
Drop table feed;
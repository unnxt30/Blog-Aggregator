-- +goose Up
CREATE Table POSTS(
    id UUID  not null,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    title varchar(100) not null,
    url varchar(100) not null UNIQUE,
    description varchar(500),
    published_at TIMESTAMP,
    feed_id UUID not null,
    PRIMARY KEY(id),
    FOREIGN KEY(feed_id) REFERENCES feed(feed_id)
);

-- +goose Down
DROP TABLE POSTS;
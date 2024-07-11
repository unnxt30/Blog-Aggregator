-- +goose Up
create table USERS(id UUID PRIMARY KEY, created_at TIMESTAMP not null, updated_at timestamp not null, name varchar(30));

-- +goose Down
drop table USERS;
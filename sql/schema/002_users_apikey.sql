-- +goose Up

alter table USERS
add api_key varchar(64) unique not null default encode(sha256(random()::text::bytea), 'hex');

-- +goose Down

Update table USERS drop column api_key;
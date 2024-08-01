-- +goose Up
alter table feed
add last_fetched_at TIMESTAMP;

-- +goose Down
update table feed drop column last_fetched_at;
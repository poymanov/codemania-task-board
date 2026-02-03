-- +goose Up
create table boards
(
    id          serial primary key,
    name        varchar   not null,
    description varchar   not null,
    owner_id    int       not null,
    created_at  timestamp not null default now(),
    updated_at  timestamp
);

-- +goose Down
drop table boards;
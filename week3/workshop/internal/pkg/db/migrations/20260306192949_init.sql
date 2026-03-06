-- +goose Up
create table articles(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL DEFAULT '',
    rating int not null default 0,
    created_at timestamp with time zone default now() not null
);

-- +goose Down
drop table articles;

-- +goose Up
-- +goose StatementBegin
create table transaction
(
    id bigserial primary key,
    name varchar(255) not null,
    cost float not null,
    time timestamp without time zone not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists transaction;
-- +goose StatementEnd

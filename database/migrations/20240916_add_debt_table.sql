-- +goose Up
-- +goose StatementBegin

create table debt
(
    id bigserial primary key,
    name varchar(100) not null,
    lender varchar(200) not null,
    borrower varchar(200) not null,
    interest_rate float,
    borrowed_amount float not null,
    paid_amount float not null,
    lend_date timestamp without time zone not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table debt;
-- +goose StatementEnd

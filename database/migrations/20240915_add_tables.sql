-- +goose Up
-- +goose StatementBegin
create table account
(
    id bigserial primary key,
    name varchar(100) not null,
    acc_balance float,
    acc_num varchar(100) not null,
    card_num varchar(100) not null,
    pin varchar(50) not null,
    security_code varchar(50) not null,
    credit_limit float,
    method_id int references method(id) not null
);
create table installment
(
    id bigserial primary key,
    name varchar(100) not null,
    total_cost float not null,
    interest_rate float not null,
    period_num int not null,
    paid_cost float not null,
    current_period int not null,
    period_cost float not null,
    account_id int references account(id) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table account;
drop table installment;
-- +goose StatementEnd

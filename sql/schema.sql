-- create table transaction
-- (
--     id bigserial primary key,
--     name varchar(255) not null,
--     cost float not null,
--     time timestamp without time zone not null default now()
-- );

create table transaction
(
    id bigserial primary key,
    name varchar(255) not null,
    cost float not null,
    time timestamp without time zone not null default now(),
    category_id int references category(id) not null
);
create table category
(
    id bigserial primary key,
    name varchar(100) not null
);
create table method
(
    id bigserial primary key,
    name varchar(100)
);
create table balancesheet
(
    id bigserial primary key,
    month int not null,
    year int not null,
    allocation float not null,
    paid float not null,
    remaining float not null,
    category_id int references category(id) not null
);
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

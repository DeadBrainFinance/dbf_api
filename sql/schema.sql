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
)

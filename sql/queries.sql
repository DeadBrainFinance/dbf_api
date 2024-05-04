-- name: CreateTransaction :one
insert into transaction (name, cost, time, category_id)
values($1, $2, $3, $4)
returning *;

-- name: GetTransaction :one
select *
from transaction
where id = $1
limit 1;

-- name: UpdateTransaction :one
update transaction
set name = $2,
    cost = $3,
    time = $4
where id = $1
returning *;

-- name: PartialUpdateTransaction :one
update transaction
set name = case when @update_name::boolean then @name::VARCHAR(255) else name end,
    cost = case when @update_cost::boolean then @cost::real else cost end,
    time = case when @update_time::boolean then @time::timestamp else time end,
    category_id = case when @update_category_id::boolean then @category_id::int else category_id end
where id = @id
returning *;

-- name: DeleteTransaction :exec
delete
from transaction
where id = $1;

-- name: ListTransactions :many
select *
from transaction
order by name;

-- name: CreateBalanceSheet :one
insert into balancesheet (month, year, allocation, paid, remaining, category_id)
values($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetBalanceSheet :one
select *
from balancesheet
where id = $1
limit 1;

-- name: PartialUpdateBalanceSheet :one
update balancesheet
set month = case when @update_month::boolean then @month::int else month end,
    year = case when @update_year::boolean then @year::int else year end,
    allocation = case when @update_allocation::boolean then @allocation::float else allocation end,
    paid = case when @update_paid::boolean then @paid::float else paid end,
    remaining = case when @update_remaining::boolean then @remaining::float else remaining end,
    category_id = case when @update_categories::boolean then @category_id::bigserial else category_id end
where id = @id
returning *;

-- name: ListBalanceSheets :many
select *
from balancesheet
order by year desc, month desc;

-- name: DeleteBalanceSheet :exec
delete
from balancesheet
where id = $1;

-- name: CreateCategory :one
insert into category (name)
values($1)
returning *;

-- name: GetCategory :one
select *
from category
where id = $1
limit 1;

-- name: PartialUpdateCategory :one
update category
set name = $2
where id = $1
returning *;

-- name: ListCategories :many
select *
from category
order by name;

-- name: DeleteCategory :exec
delete
from category
where id = $1;

-- name: CreateTransaction :one
insert into transaction (name, cost, time)
values($1, $2, $3)
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
    time = case when @update_time::boolean then @time::timestamp else time end
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

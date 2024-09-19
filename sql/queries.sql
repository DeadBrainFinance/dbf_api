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

-- name: CreateAccount :one
insert into account (name, acc_balance, acc_num, card_num, pin, security_code, credit_limit, method_id)
values($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: GetAccount :one
select *
from account
where id = $1
limit 1;

-- name: PartialUpdateAccount :one
update account
set name = case when @update_name::boolean then @name::varchar(100) else name end,
    acc_balance = case when @update_acc_balance::boolean then @acc_balance::int else acc_balance end,
    acc_num = case when @update_acc_num::boolean then @acc_num::varchar(100) else acc_num end,
    card_num = case when @update_card_num::boolean then @card_num::varchar(100) else card_num end,
    pin = case when @update_pin::boolean then @pin::varchar(50) else pin end,
    security_code = case when @update_security_code::boolean then @security_code::varchar(50) else security_code end,
    credit_limit = case when @update_credit_limit::boolean then @credit_limit::float else credit_limit end,
    method_id = case when @update_method::boolean then @method_id::bigserial else method_id end
where id = @id
returning *;

-- name: ListAccounts :many
select *
from account
order by name;

-- name: DeleteAccount :exec
delete
from account
where id = $1;

-- name: CreateInstallment :one
insert into installment (name, total_cost, interest_rate, period_num, paid_cost, current_period, period_cost, account_id)
values($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: GetInstallment :one
select *
from installment
where id = $1
limit 1;

-- name: PartialUpdateInstallment :one
update installment
set installment_name = case when @update_name::boolean then @name::varchar(100) else name end,
    total_cost = case when @update_total_cost::boolean then @total_cost::float else total_cost end,
    interest_rate = case when @update_interest_rate::boolean then @interest_rate::float else interest_rate end,
    period_num = case when @update_period_num::boolean then @period_num::int else period_num end,
    paid_cost = case when @update_paid_cost::boolean then @paid_cost::float else paid_cost end,
    current_period = case when @update_current_period::boolean then @current_period::int else current_period end,
    period_cost = case when @update_period_cost::boolean then @period_cost::float else period_cost end,
    account_id = case when @update_account::boolean then @account_id::bigserial else account_id end
where id = @id
returning *;

-- name: ListInstallments :many
select *
from installment
order by name;

-- name: DeleteInstallment :exec
delete
from installment
where id = $1;


-- name: CreateDebt :one
insert into debt (name, lender, borrower, interest_rate, borrowed_amount, paid_amount, lend_date)
values($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: GetDebt :one
select *
from debt
where id = $1
limit 1;


-- name: PartialUpdateDebt :one
update debt
set name = case when @update_name::boolean then @name::varchar(100) else name end,
    lender = case when @update_lender::boolean then @lender::varchar(200) else lender end,
    borrower = case when @update_borrower::boolean then @borrower::varchar(200) else borrower end,
    interest_rate = case when @update_interest_rate::boolean then @interest_rate::float else interest_rate end,
    paid_amount = case when @update_paid_amount::boolean then @paid_amount::float else paid_amount end,
    lend_date = case when @update_lend_date::boolean then @lend_date::timestamp else lend_date end,
where id = @id
returning *;

-- name: ListDebts :many
select *
from debt
order by name;

-- name: DeleteDebt :exec
delete
from debt
where id = $1;

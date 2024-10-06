package utils

const GetInstallment = `-- name: GetInstallment :one
select id, name, total_cost, interest_rate, period_num, paid_cost, current_period, period_cost, account_id
from installment
where id = $1
limit 1
`
const PartialUpdateInstallment = `-- name: PartialUpdateInstallment :one
update installment
set name = case when $1::boolean then $2::varchar(100) else name end,
    total_cost = case when $3::boolean then $4::float else total_cost end,
    interest_rate = case when $5::boolean then $6::float else interest_rate end,
    period_num = case when $7::boolean then $8::int else period_num end,
    paid_cost = case when $9::boolean then $10::float else paid_cost end,
    current_period = case when $11::boolean then $12::int else current_period end,
    period_cost = case when $13::boolean then $14::float else period_cost end,
    account_id = case when $15::boolean then $16::int else account_id end
where id = $17
returning id, name, total_cost, interest_rate, period_num, paid_cost, current_period, period_cost, account_id
`
const DeleteInstallment = `-- name: DeleteInstallment :exec
delete
from installment
where id = $1
`
const ListInstallments = `-- name: ListInstallments :many
select id, name, total_cost, interest_rate, period_num, paid_cost, current_period, period_cost, account_id
from installment
order by name
`
const CreateInstallment = `-- name: CreateInstallment :one
insert into installment (name, total_cost, interest_rate, period_num, paid_cost, current_period, period_cost, account_id)
values($1, $2, $3, $4, $5, $6, $7, $8)
returning id, name, total_cost, interest_rate, period_num, paid_cost, current_period, period_cost, account_id
`

const GetAccount = `-- name: GetAccount :one
select id, name, coalesce(acc_balance, 0), acc_num, card_num, pin, security_code, coalesce(credit_limit, 0), method_id
from account
where id = $1
limit 1
`
const CreateAccount = `-- name: CreateAccount :one
insert into account (name, acc_balance, acc_num, card_num, pin, security_code, credit_limit, method_id)
values($1, $2, $3, $4, $5, $6, $7, $8)
returning id, name, acc_balance, acc_num, card_num, pin, security_code, credit_limit, method_id
`
const DeleteAccount = `-- name: DeleteAccount :exec
delete
from account
where id = $1
`
const PartialUpdateAccount = `-- name: PartialUpdateAccount :one
update account
set name = case when $1::boolean then $2::varchar(100) else name end,
    acc_balance = case when $3::boolean then $4::int else acc_balance end,
    acc_num = case when $5::boolean then $6::varchar(100) else acc_num end,
    card_num = case when $7::boolean then $8::varchar(100) else card_num end,
    pin = case when $9::boolean then $10::varchar(50) else pin end,
    security_code = case when $11::boolean then $12::varchar(50) else security_code end,
    credit_limit = case when $13::boolean then $14::float else credit_limit end,
    method_id = case when $15::boolean then $16::int else method_id end
where id = $17
returning id, name, acc_balance, acc_num, card_num, pin, security_code, credit_limit, method_id
`
const ListAccounts = `-- name: ListAccounts :many
select id, name, coalesce(acc_balance, 0), acc_num, card_num, pin, security_code, coalesce(credit_limit, 0), method_id
from account
order by name
`

const GetBalanceSheet = `-- name: GetBalanceSheet :one
select id, month, year, allocation, paid, remaining, category_id
from balancesheet
where id = $1
limit 1
`
const CreateBalanceSheet = `-- name: CreateBalanceSheet :one
insert into balancesheet (month, year, allocation, paid, remaining, category_id)
values($1, $2, $3, $4, $5, $6)
returning id, month, year, allocation, paid, remaining, category_id
`
const DeleteBalanceSheet = `-- name: DeleteBalanceSheet :exec
delete
from balancesheet
where id = $1
`
const PartialUpdateBalanceSheet = `-- name: PartialUpdateBalanceSheet :one
update balancesheet
set month = case when $1::boolean then $2::int else month end,
    year = case when $3::boolean then $4::int else year end,
    allocation = case when $5::boolean then $6::real else allocation end,
    paid = case when $7::boolean then $8::real else paid end,
    remaining = case when $9::boolean then $10::real else remaining end,
    category_id = case when $11::boolean then $12::int else category_id end
where id = $13
returning id, month, year, allocation, paid, remaining, category_id
`
const ListBalanceSheets = `-- name: ListBalanceSheets :many
select id, month, year, allocation, paid, remaining, category_id
from balancesheet
order by year desc, month desc
`
const CreateMethod = `-- name: CreateMethod :one
insert into method (name)
values($1)
returning id, name
`
const GetMethod = `-- name: GetMethod :one
select id, name
from method
where id = $1
limit 1
`
const PartialUpdateMethod = `-- name: PartialUpdateMethod :one
update method
set name = $2
where id = $1
returning id, name
`
const ListMethods = `-- name: ListMethods :many
select id, coalesce(name, '')
from method
order by name
`
const DeleteMethod = `-- name: DeleteMethod :exec
delete
from method
where id = $1
`

const GetTransaction = `-- name: GetTransaction :one
select id, name, cost, time, category_id
from transaction
where id = $1
limit 1
`
const CreateTransaction = `-- name: CreateTransaction :one
insert into transaction (name, cost, time, category_id)
values($1, $2, $3, $4)
returning id, name, cost, time, category_id
`
const DeleteTransaction = `-- name: DeleteTransaction :exec
delete
from transaction
where id = $1
`
const PartialUpdateTransaction = `-- name: PartialUpdateTransaction :one
update transaction
set name = case when $1::boolean then $2::varchar(255) else name end,
    cost = case when $3::boolean then $4::real else cost end,
    time = case when $5::boolean then $6::timestamp else time end,
    category_id = case when $7::boolean then $8::int else category_id end
where id = $9
returning id, name, cost, time, category_id
`
const ListTransactions = `-- name: ListTransactions :many
select id, name, cost, time, category_id
from transaction
order by name
`


const CreateDebt = `-- name: CreateDebt :one
insert into debt (name, lender, borrower, interest_rate, borrowed_amount, paid_amount, lend_date)
values($1, $2, $3, $4, $5, $6, $7)
returning id, name, lender, borrower, interest_rate, borrowed_amount, paid_amount, lend_date
`
const GetDebt = `-- name: GetDebt :one
select id, name, lender, borrower, coalesce(interest_rate, 0), borrowed_amount, paid_amount, lend_date
from debt
where id = $1
limit 1
`
const PartialUpdateDebt = `-- name: PartialUpdateDebt :one
update debt
set name = case when $1::boolean then $2::varchar(100) else name end,
    lender = case when $3::boolean then $4::varchar(200) else  lender end,
    borrower = case when $5::boolean then $6::varchar(200) else borrower end,
    interest_rate = case when $7::boolean then $8::float else interest_rate end,
    paid_amount = case when $9::boolean then $10::float else paid_amount end,
    lend_date = case when $11::boolean then $12::timestamp else lend_date end,
where id = $13
returning id, name, lender, borrower, interest_rate, borrowed_amount, paid_amount, lend_date
`
const ListDebts = `-- name: ListDebts :many
select id, name, lender, borrower, coalesce(interest_rate, 0), borrowed_amount, paid_amount, lend_date
from debt
order by name
`
const DeleteDebt = `-- name: DeleteDebt :exec
delete
from debt
where id = $1
`


const CreateCategory = `-- name: CreateCategory :one
insert into category (name, description)
values($1, $2)
returning id, name, description
`
const GetCategory = `-- name: GetCategory :one
select id, name, coalesce(description, '')
from category
where id = $1
limit 1
`
const PartialUpdateCategory = `-- name: PartialUpdateCategory :one
update category
set name = case when $1:boolean then $2::varchar(100) else name end,
    description = case when $3::boolean then $4::varchar(200) else description end,
where id = $5
returning id, name, description
`
const ListCategories = `-- name: ListCategories :many
select id, name, coalesce(description, '')
from category
order by name
`
const DeleteCategory = `-- name: DeleteCategory :exec
delete
from category
where id = $1
`
